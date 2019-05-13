import * as net from 'net';

export class Connection {
  private readonly server: string;
  private readonly port: number;
  private socket: net.Socket;

  constructor(server: string, port: number, errorHandler: (error: Error) => void) {
    this.server = server;
    this.port = port;
    this.socket = new net.Socket();
    this.socket.on('error', errorHandler);
  }

  initiateHandshake() {
    this.socket.connect(this.port, this.server);
    this.socket.addListener('data', (e) => {
      console.log('received data');
      console.dir(e);
      console.dir(new TextDecoder('utf-8').decode(e));

      this.socket.write((new TextEncoder()).encode('Hello from client\n'));
    });
  }
}
