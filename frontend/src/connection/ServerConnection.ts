import { Connection } from './Connection';
import { GameList, User } from './proto/Types_pb';
import { Protocol } from './protocol/Commands';
import { Message } from './Message';

export class ServerConnection extends Connection {
  async sendConnect(userName: string): Promise<Message<GameList>> {
    const user = new User();
    user.setName(userName);

    return await this.send<GameList>(new Message(Protocol.ClientInstructions.CONNECTION.CONNECT, user));
  }
}
