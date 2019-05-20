import { Protocol } from './protocol/Commands';
import * as Protobuf from 'google-protobuf';
import { ProtoPayload } from './proto/ProtoPayload';

export class Message {
  command: Protocol.Command;
  payload?: Protobuf.Message;

  constructor(command: Protocol.Command, payload?: ProtoPayload) {
    this.command = command;
    this.payload = payload;
  }
}
