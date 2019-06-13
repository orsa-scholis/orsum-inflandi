import * as net from 'net';
import { Message } from './Message';
import PacketQueue from './PacketQueue';
import Packet from './Packet';
import PacketSerializer from './PacketSerializer';
import { ProtoPayload } from './proto/ProtoPayload';
import { ResolverCallback } from './EnqueuedPacket';

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

  send<T extends ProtoPayload>(message: Message<T>): Promise<Message<T>> {
    return new Promise<Message<T>>((resolve, reject) => {
      const packet = new Packet(message);
      this.queue.enqueue(packet, resolve as ResolverCallback, reject);
      this.transmitPacket(packet);
    });
  }

  private transmitPacket<T extends ProtoPayload>(packet: Packet<T>) {
    this.socket.write(new PacketSerializer(packet).serialize());
  }
}
