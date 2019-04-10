import * as net from 'net';

class Connection {
  private readonly server: string;
  private readonly port: number;
  private socket: net.Socket;

  constructor(server: string, port: number, errorHandler: () => void) {
    this.server = server;
    this.port = port;
    this.socket = new net.Socket();
    this.socket.on('error', errorHandler);
  }

  initiateHandshake() {
    this.socket.connect(this.port, this.server);
  }
}
