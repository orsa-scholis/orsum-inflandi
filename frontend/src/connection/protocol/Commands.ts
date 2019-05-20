export namespace Protocol {
  class ProtocolCommand {
    readonly domain: string;
    readonly command: string;

    constructor(domain: string, command: string) {
      this.domain = domain;
      this.command = command;
    }

    toString() {
      return `${this.domain}:${this.command}`;
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
}
