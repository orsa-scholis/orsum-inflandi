import * as net from 'net';

export class Connection {
  private readonly server: string;
  private readonly port: number;
  private socket: net.Socket;

  constructor(server: string, port: number, errorHandler: () => void) {
    this.server = server;
    this.port = port;
    this.socket = new net.Socket();
    this.socket.on('error', errorHandler);

    this.socket.on('data', (data) => {
      console.log(new TextDecoder('utf-8').decode(data));
    });

    this.initiateHandshake();
  }

  initiateHandshake() {
    this.socket.connect(this.port, this.server);
  }
}
