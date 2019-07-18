import { Message } from './Message';
import * as uuid from 'uuid/v1';
import { ProtoPayload } from './proto/ProtoPayload';

export default class Packet<T extends ProtoPayload> {
  readonly message: Message<T>;
  readonly uuid: string;

  constructor(message: Message<T>, packetUuid: string = uuid()) {
    this.message = message;
    this.uuid = packetUuid;
  }
}
