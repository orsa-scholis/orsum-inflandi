import { Message } from './Message';
import * as uuid from 'uuid/v1';

export default class Packet {
  readonly message: Message;
  readonly uuid: string;

  constructor(message: Message, packetUuid: string = uuid()) {
    this.message = message;
    this.uuid = packetUuid;
  }
}
