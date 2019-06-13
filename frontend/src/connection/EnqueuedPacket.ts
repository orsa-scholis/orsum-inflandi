import Packet from './Packet';
import { Message } from './Message';
import { ProtoPayload } from './proto/ProtoPayload';

export type ResolverCallback = <T extends ProtoPayload>(value?: Message<T> | PromiseLike<Message<T>>) => void;
export type RejectionCallback = (reason?: any) => void;

export class EnqueuedPacket<T extends ProtoPayload> {
  packet: Packet<T>;
  private readonly resolverCallback: ResolverCallback;
  private readonly rejectionCallback: RejectionCallback;

  constructor(packet: Packet<T>, resolverCallback: ResolverCallback, rejectionCallback: RejectionCallback) {
    this.packet = packet;
    this.resolverCallback = resolverCallback;
    this.rejectionCallback = rejectionCallback;
  }

  resolve(message: Message<T>) {
    this.resolverCallback(message);
  }

  reject(reason: any) {
    this.rejectionCallback(reason);
  }
}
