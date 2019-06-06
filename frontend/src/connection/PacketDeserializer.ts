import Packet from './Packet';
import SerializationError from './SerializationError';
import { Message } from './Message';
import { Protocol } from './protocol/Commands';
import { Message as ProtobufMessage } from 'google-protobuf';

const UUID_MATCHER = Object.freeze(/(?:(?:\w|\d){4,}-){4}(?:\w|\d){12}/);
const ALL_SERVER_COMMANDS = Object.values(Protocol.PossibleServerCommands).reduce((r: Protocol.Command[], v) => {
  return r.concat((v.constructor.name === 'ProtocolCommand') ? [v] : Object.values(v));
}, []);

interface ProtobufMessageClass<T> {
  new(...args: any[]): T;
  deserializeBinary(bytes: Uint8Array): ProtobufMessage;
}

export default class PacketDeserializer<PayloadType extends ProtobufMessage> {
  private input: string;
  private raw: {
    uuid: string;
    domain: string;
    command: string;
    payload: string;
  };
  private payloadClass?: ProtobufMessageClass<PayloadType>;

  constructor(input: string, payloadClass?: ProtobufMessageClass<PayloadType>) {
    this.input = input;
    this.payloadClass = payloadClass;
  }

  deserialize(): Packet {
    const splitted = this.input.split(':');
    if (splitted.length < 3) {
      throw new SerializationError('Message has invalid segmentation');
    }

    const [uuid, domain, command, payload] = splitted;
    this.raw = { uuid, domain, command, payload };

    return new Packet(this.deserializeMessage(), this.deserializeUUID());
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

  private deserializeMessage(): Message {
    const needle = new Protocol.ProtocolCommand(this.raw.domain, this.raw.command).toString();
    const messageCommand = ALL_SERVER_COMMANDS.find(command => command.toString() === needle);

    if (messageCommand === undefined) {
      throw new SerializationError(`Unknown command '${needle}'`);
    }

    return new Message(messageCommand, this.deserializePayload());
  }

  private deserializePayload(): PayloadType | undefined {
    if (!this.raw.payload || !this.payloadClass) {
      return;
    }

    return <PayloadType>this.payloadClass.deserializeBinary(Buffer.from(this.raw.payload, 'base64'));
  }
}
