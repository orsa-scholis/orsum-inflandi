import ProtocolInstruction from './ProtocolInstruction';
import ClientProtocolInstruction from './ClientProtocolInstruction';
import { Game, GameList } from '../proto/Types_pb';
import ProtobufMessageClass from '../ProtobufMessageClass';

export namespace Protocol {
  export type Instruction = ProtocolInstruction | ClientProtocolInstruction<ProtobufMessageClass<any>>;

  export const ClientInstructions = Object.freeze({
    CONNECTION: {
      CONNECT: new ClientProtocolInstruction('connection', 'connect', GameList)
    },
    GAME: {
      NEW: new ClientProtocolInstruction('game', 'new', Game),
      JOIN: new ClientProtocolInstruction('game', 'join'),
      TURN: new ClientProtocolInstruction('game', 'turn')
    },
    CHAT: {
      SEND: new ClientProtocolInstruction('chat', 'send')
    },
  });

  export const PossibleServerInstructions = Object.freeze({
    SUCCESS: new ProtocolInstruction('success'),
    FAILURE: new ProtocolInstruction('failure'),
    BROADCAST: {
      CHAT: new ProtocolInstruction('broadcast', 'chat'),
      GAMES: new ProtocolInstruction('broadcast', 'games'),
      TURN: new ProtocolInstruction('broadcast', 'turn'),
      END: new ProtocolInstruction('broadcast', 'end')
    }
  });
}
