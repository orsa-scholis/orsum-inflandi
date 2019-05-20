import Packet from './Packet';

export default class PacketQueue {
  private queue: Packet[] = [];

  enqueue(packet: Packet) {
    this.queue.push(packet);
  }

  findPacket(packetUUID: string) {
    return this.queue.find(packet => packet.uuid == packetUUID);
  }

  popPacket(answerUUID: string) {
    const packet = this.findPacket(answerUUID);
    if (!packet) {
      return undefined;
    }

    const index = this.queue.indexOf(packet);
    this.queue.splice(index, 1);
    return packet;
  }
}
