import { Message } from './Message';
import * as uuid from 'uuid/v1';

type ResolverCallback = <T extends Message>(value?: T | PromiseLike<T>) => void;
type RejectionCallback = (reason?: any) => void;

const MESSAGE_DELIMITER = (new TextEncoder()).encode('\n');

export default class Packet {
  readonly message: Message;
  readonly uuid: string = uuid();
  readonly resolver: ResolverCallback;
  readonly rejectionCallback: RejectionCallback;

  constructor(message: Message, resolver: ResolverCallback, rejectionCallback: RejectionCallback) {
    this.message = message;
    this.resolver = resolver;
    this.rejectionCallback = rejectionCallback;
  }

  serializeMessage(): Uint8Array {
    const uuidContent = new TextEncoder().encode(`${this.uuid}:`);
    const messageContent = Packet.appendToUint8Array(this.serializeCommand(), this.serializePayload());

    return Packet.appendToUint8Array(uuidContent, Packet.appendToUint8Array(messageContent, MESSAGE_DELIMITER));
  }

  private serializePayload() {
    if (!this.message.payload) {
      return undefined;
    }

    const base64 = Buffer.from(this.message.payload.serializeBinary()).toString('base64');
    return (new TextEncoder()).encode(`:${base64}`);
  }

  private serializeCommand() {
    return (new TextEncoder()).encode(this.message.command.toString());
  }

  private static appendToUint8Array(left: Uint8Array, right?: Uint8Array) {
    if (!right) {
      return left;
    }

    const output = new Uint8Array(left.length + right.length);
    output.set(left, 0);
    output.set(right, left.length);

    return output;
  }
}
