import { Protocol } from './protocol/Commands';
import { ProtoPayload } from './proto/ProtoPayload';

export class Message {
  command: Protocol.Command;
  payload?: ProtoPayload;

  constructor(command: Protocol.Command, payload?: ProtoPayload) {
    this.command = command;
    this.payload = payload;
  }
}
