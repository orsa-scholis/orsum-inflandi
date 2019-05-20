import {
  ChatMessage,
  FourInARowTurnPayload,
  Game,
  GameEnd,
  GameList,
  GameRequest,
  User
} from './Types_pb';

export type ProtoPayload = User | GameRequest | Game | GameList | FourInARowTurnPayload | ChatMessage | GameEnd;
