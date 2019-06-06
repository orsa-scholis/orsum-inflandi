export namespace Protocol {
  export class ProtocolCommand {
    readonly domain: string;
    readonly command?: string;

    constructor(domain: string, command?: string) {
      this.domain = domain;
      this.command = command;
    }

    toString() {
      return this.command ? `${this.domain}:${this.command}` : this.domain;
    }
  }

  export type Command = ProtocolCommand;

  export const ClientCommands = Object.freeze({
    CONNECTION: {
      CONNECT: new ProtocolCommand('connection', 'connect')
    },
    GAME: {
      NEW: new ProtocolCommand('game', 'new'),
      JOIN: new ProtocolCommand('game', 'join'),
      TURN: new ProtocolCommand('game', 'turn')
    },
    CHAT: {
      SEND: new ProtocolCommand('chat', 'send')
    },
  });

  export const PossibleServerCommands = Object.freeze({
    SUCCESS: new ProtocolCommand('success'),
    FAILURE: new ProtocolCommand('failure'),
    BROADCAST: {
      CHAT: new ProtocolCommand('broadcast', 'chat'),
      GAMES: new ProtocolCommand('broadcast', 'games'),
      TURN: new ProtocolCommand('broadcast', 'turn'),
      END: new ProtocolCommand('broadcast', 'end')
    }
  });
}
