import Packet from './Packet';

const MESSAGE_DELIMITER = Object.seal((new TextEncoder()).encode('\n'));

function appendToUint8Array(left: Uint8Array, right?: Uint8Array) {
  if (!right) {
    return left;
  }

  const output = new Uint8Array(left.length + right.length);
  output.set(left, 0);
  output.set(right, left.length);

  return output;
}

export default class PacketSerializer {
  private packet: Packet;

  constructor(packet: Packet) {
    this.packet = packet;
  }

  serialize(): Uint8Array {
    const uuidContent = new TextEncoder().encode(`${this.packet.uuid}:`);
    const messageContent = appendToUint8Array(this.serializeCommand(), this.serializePayload());

    return appendToUint8Array(uuidContent, appendToUint8Array(messageContent, MESSAGE_DELIMITER));
  }

  private serializePayload() {
    if (!this.packet.message.payload) {
      return undefined;
    }

    const base64 = Buffer.from(this.packet.message.payload.serializeBinary()).toString('base64');
    return (new TextEncoder()).encode(`:${base64}`);
  }

  private serializeCommand() {
    return (new TextEncoder()).encode(this.packet.message.instruction.toString());
  }
}
