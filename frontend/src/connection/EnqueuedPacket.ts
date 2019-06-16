import Packet from './Packet';
import { Message } from './Message';
import { ProtoPayload } from './proto/ProtoPayload';

export type ResolverCallback = <T extends ProtoPayload>(value?: Message<T> | PromiseLike<Message<T>>) => void;
export type RejectionCallback = (reason?: any) => void;

export class EnqueuedPacket {
  packet: Packet<ProtoPayload>;
  private readonly resolverCallback: ResolverCallback;
  private readonly rejectionCallback: RejectionCallback;

  constructor(packet: Packet<ProtoPayload>, resolverCallback: ResolverCallback, rejectionCallback: RejectionCallback) {
    this.packet = packet;
    this.resolverCallback = resolverCallback;
    this.rejectionCallback = rejectionCallback;
  }

  resolve(message: Message<ProtoPayload>) {
    this.resolverCallback(message);
  }

  reject(reason: any) {
    this.rejectionCallback(reason);
  }
}
