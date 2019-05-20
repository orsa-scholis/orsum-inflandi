import * as net from 'net';
import { Message } from './Message';
import PacketQueue from './PacketQueue';
import Packet from './Packet';

export class Connection {
  private readonly server: string;
  private readonly port: number;
  private readonly queue: PacketQueue = new PacketQueue();
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

      // TODO: Deserialize and resolve/reject
    });
  }

  send(message: Message): Promise<Message> {
    return new Promise<Message>((resolve, reject) => {
      const packet = new Packet(message, resolve, reject);
      this.queue.enqueue(packet);
      this.transmitPacket(packet);
    });
  }

  private transmitPacket(packet: Packet) {
    this.socket.write(packet.serializeMessage());
  }
}
