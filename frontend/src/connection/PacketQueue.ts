import Packet from './Packet';
import { EnqueuedPacket, RejectionCallback, ResolverCallback } from './EnqueuedPacket';
import { ProtoPayload } from './proto/ProtoPayload';

export default class PacketQueue {
  private queue: EnqueuedPacket[] = [];

  enqueue(packet: Packet<ProtoPayload>, resolverCallback: ResolverCallback, rejectionCallback: RejectionCallback) {
    this.queue.push(new EnqueuedPacket(packet, resolverCallback, rejectionCallback));
  }

  findPacket(packetUUID: string) {
    return this.queue.find(enqueuedPacket => enqueuedPacket.packet.uuid == packetUUID);
  }

  popPacket(answerUUID: string): EnqueuedPacket | undefined {
    const packet = this.findPacket(answerUUID);
    if (!packet) {
      return undefined;
    }

    return this.queue.splice(this.queue.indexOf(packet), 1)[0];
  }
}
