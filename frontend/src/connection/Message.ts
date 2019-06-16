import { Protocol } from './protocol/Commands';
import { ProtoPayload } from './proto/ProtoPayload';

export class Message<T extends ProtoPayload> {
  instruction: Protocol.Instruction;
  payload?: T;

  constructor(command: Protocol.Instruction, payload?: T) {
    this.instruction = command;
    this.payload = payload;
  }
}
