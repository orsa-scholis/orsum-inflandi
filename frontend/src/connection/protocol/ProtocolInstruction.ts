export default class ProtocolInstruction {
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
