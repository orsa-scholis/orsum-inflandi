import Packet from './Packet';
import { Message } from './Message';

export type ResolverCallback = <T extends Message>(value?: T | PromiseLike<T>) => void;
export type RejectionCallback = (reason?: any) => void;

export class EnqueuedPacket {
  packet: Packet;
  private readonly resolverCallback: ResolverCallback;
  private readonly rejectionCallback: RejectionCallback;

  constructor(packet: Packet, resolverCallback: ResolverCallback, rejectionCallback: RejectionCallback) {
    this.packet = packet;
    this.resolverCallback = resolverCallback;
    this.rejectionCallback = rejectionCallback;
  }

  resolve(message: Message) {
    this.resolverCallback(message);
  }

  reject(reason: any) {
    this.rejectionCallback(reason);
  }
}
