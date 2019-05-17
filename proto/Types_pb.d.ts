// package: 
// file: proto/Types.proto

import * as jspb from "google-protobuf";

export class GameRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getType(): GameTypeMap[keyof GameTypeMap];
  setType(value: GameTypeMap[keyof GameTypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GameRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GameRequest): GameRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GameRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GameRequest;
  static deserializeBinaryFromReader(message: GameRequest, reader: jspb.BinaryReader): GameRequest;
}

export namespace GameRequest {
  export type AsObject = {
    name: string,
    type: GameTypeMap[keyof GameTypeMap],
  }
}

export class User extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    name: string,
  }
}

export class Game extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getType(): GameTypeMap[keyof GameTypeMap];
  setType(value: GameTypeMap[keyof GameTypeMap]): void;

  getId(): number;
  setId(value: number): void;

  hasInitiator(): boolean;
  clearInitiator(): void;
  getInitiator(): User | undefined;
  setInitiator(value?: User): void;

  hasOpponent(): boolean;
  clearOpponent(): void;
  getOpponent(): User | undefined;
  setOpponent(value?: User): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Game.AsObject;
  static toObject(includeInstance: boolean, msg: Game): Game.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Game, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Game;
  static deserializeBinaryFromReader(message: Game, reader: jspb.BinaryReader): Game;
}

export namespace Game {
  export type AsObject = {
    name: string,
    type: GameTypeMap[keyof GameTypeMap],
    id: number,
    initiator?: User.AsObject,
    opponent?: User.AsObject,
  }
}

export class GameList extends jspb.Message {
  clearGamesList(): void;
  getGamesList(): Array<Game>;
  setGamesList(value: Array<Game>): void;
  addGames(value?: Game, index?: number): Game;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GameList.AsObject;
  static toObject(includeInstance: boolean, msg: GameList): GameList.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GameList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GameList;
  static deserializeBinaryFromReader(message: GameList, reader: jspb.BinaryReader): GameList;
}

export namespace GameList {
  export type AsObject = {
    gamesList: Array<Game.AsObject>,
  }
}

export class FourInARowTurnPayload extends jspb.Message {
  getGameid(): number;
  setGameid(value: number): void;

  getRow(): number;
  setRow(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FourInARowTurnPayload.AsObject;
  static toObject(includeInstance: boolean, msg: FourInARowTurnPayload): FourInARowTurnPayload.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: FourInARowTurnPayload, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FourInARowTurnPayload;
  static deserializeBinaryFromReader(message: FourInARowTurnPayload, reader: jspb.BinaryReader): FourInARowTurnPayload;
}

export namespace FourInARowTurnPayload {
  export type AsObject = {
    gameid: number,
    row: number,
  }
}

export class ChatMessage extends jspb.Message {
  getContext(): ChatMessageContextMap[keyof ChatMessageContextMap];
  setContext(value: ChatMessageContextMap[keyof ChatMessageContextMap]): void;

  hasUser(): boolean;
  clearUser(): void;
  getUser(): User | undefined;
  setUser(value?: User): void;

  getContent(): string;
  setContent(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChatMessage.AsObject;
  static toObject(includeInstance: boolean, msg: ChatMessage): ChatMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ChatMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChatMessage;
  static deserializeBinaryFromReader(message: ChatMessage, reader: jspb.BinaryReader): ChatMessage;
}

export namespace ChatMessage {
  export type AsObject = {
    context: ChatMessageContextMap[keyof ChatMessageContextMap],
    user?: User.AsObject,
    content: string,
  }
}

export class GameEnd extends jspb.Message {
  getResult(): GameResultMap[keyof GameResultMap];
  setResult(value: GameResultMap[keyof GameResultMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GameEnd.AsObject;
  static toObject(includeInstance: boolean, msg: GameEnd): GameEnd.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GameEnd, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GameEnd;
  static deserializeBinaryFromReader(message: GameEnd, reader: jspb.BinaryReader): GameEnd;
}

export namespace GameEnd {
  export type AsObject = {
    result: GameResultMap[keyof GameResultMap],
  }
}

export interface GameTypeMap {
  FOUR_IN_A_ROW: 0;
}

export const GameType: GameTypeMap;

export interface ChatMessageContextMap {
  LOBBY: 0;
  IN_GAME: 1;
}

export const ChatMessageContext: ChatMessageContextMap;

export interface GameResultMap {
  WON: 0;
  LOST: 1;
  TIE: 2;
}

export const GameResult: GameResultMap;

