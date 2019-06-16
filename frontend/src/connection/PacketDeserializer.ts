import Packet from './Packet';
import SerializationError from './SerializationError';
import { Message } from './Message';
import { Protocol } from './protocol/Commands';
import { Message as ProtobufMessage } from 'google-protobuf';
import ProtobufMessageClass from './ProtobufMessageClass';
import ProtocolInstruction from './protocol/ProtocolInstruction';

const UUID_MATCHER = Object.freeze(/(?:(?:\w|\d){4,}-){4}(?:\w|\d){12}/);
const ALL_SERVER_COMMANDS = Object.values(Protocol.PossibleServerInstructions).reduce((r: Protocol.Instruction[], v) => {
  return r.concat((v.constructor.name === 'ProtocolInstruction') ? [v] : Object.values(v));
}, []);

export default class PacketDeserializer<PayloadType extends ProtobufMessage> {
  private payloadClass?: ProtobufMessageClass<PayloadType>;
  private input: string;
  private raw: {
    uuid: string;
    domain: string;
    command: string;
    payload: string;
  };

  constructor(input: string, payloadClass?: ProtobufMessageClass<PayloadType>) {
    this.input = input;
    this.payloadClass = payloadClass;
  }

  for(payloadClass: ProtobufMessageClass<PayloadType>): this {
    this.payloadClass = payloadClass;
    return this;
  }

  deserialize(): Packet<PayloadType> {
    this.splitMessage();

    return new Packet(this.deserializeMessage(), this.deserializeUUID());
  }

  scanUUID() {
    this.splitMessage();
    return this.deserializeUUID();
  }

  private splitMessage() {
    const splitted = this.input.trim().split(':');
    if (splitted.length < 3) {
      throw new SerializationError('Message has invalid segmentation');
    }

    const [uuid, domain, command, payload] = splitted;
    this.raw = { uuid, domain, command, payload };
  }

  private deserializeUUID(): string {
    if (!this.raw.uuid.length) {
      throw new SerializationError('UUID is not present');
    }

    if (!UUID_MATCHER.test(this.raw.uuid)) {
      throw new SerializationError('UUID does not conform to RFC4122');
    }

    return this.raw.uuid;
  }

  private deserializeMessage(): Message<PayloadType> {
    return new Message(this.findMessageCommand(), this.deserializePayload());
  }

  private findMessageCommand() {
    let messageCommand = this.findServerCommand(new ProtocolInstruction(this.raw.domain, this.raw.command));

    if (messageCommand === undefined) {
      messageCommand = this.findServerCommand(new ProtocolInstruction(this.raw.domain));

      if (messageCommand === undefined) {
        throw new SerializationError(`Unknown command '${this.raw.command}' for domain '${this.raw.domain}'`);
      } else {
        this.raw.payload = this.raw.command;
      }
    }

    return messageCommand;
  }

  private findServerCommand(command: ProtocolInstruction) {
    const needle = command.toString();
    return ALL_SERVER_COMMANDS.find(command => command.toString() === needle);
  }

  private deserializePayload(): PayloadType | undefined {
    if (!this.raw.payload || !this.payloadClass) {
      return;
    }

    return <PayloadType>this.payloadClass.deserializeBinary(Buffer.from(this.raw.payload, 'base64'));
  }
}
