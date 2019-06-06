import Packet from './Packet';
import { EnqueuedPacket, RejectionCallback, ResolverCallback } from './EnqueuedPacket';

export default class PacketQueue {
  private queue: EnqueuedPacket[] = [];

  enqueue(packet: Packet, resolverCallback: ResolverCallback, rejectionCallback: RejectionCallback) {
    this.queue.push(new EnqueuedPacket(packet, resolverCallback, rejectionCallback));
  }

  findPacket(packetUUID: string) {
    return this.queue.find(enqueuedPacket => enqueuedPacket.packet.uuid == packetUUID);
  }

  popPacket(answerUUID: string) {
    const packet = this.findPacket(answerUUID);
    if (!packet) {
      return undefined;
    }

    const index = this.queue.indexOf(packet);
    return this.queue.splice(index, 1);
  }
}
