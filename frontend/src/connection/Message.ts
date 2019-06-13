import { Protocol } from './protocol/Commands';
import { ProtoPayload } from './proto/ProtoPayload';

export class Message<T extends ProtoPayload> {
  command: Protocol.Command;
  payload?: T;

  constructor(command: Protocol.Command, payload?: T) {
    this.command = command;
    this.payload = payload;
  }
}
