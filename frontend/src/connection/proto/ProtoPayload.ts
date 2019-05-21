import {
  ChatMessage,
  FourInARowTurnPayload,
  Game,
  GameEnd,
  GameList,
  GameRequest,
  User
} from './Types_pb';

export type Turn = FourInARowTurnPayload;

export type ProtoPayload = User | GameRequest | Game | GameList | Turn | ChatMessage | GameEnd;
