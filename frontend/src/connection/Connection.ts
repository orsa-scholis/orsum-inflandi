import * as net from 'net';
import { Message } from './Message';
import PacketQueue from './PacketQueue';
import Packet from './Packet';
import PacketSerializer from './PacketSerializer';
import { ProtoPayload } from './proto/ProtoPayload';
import { EnqueuedPacket, ResolverCallback } from './EnqueuedPacket';
import PacketDeserializer from './PacketDeserializer';
import ClientProtocolInstruction from './protocol/ClientProtocolInstruction';
import { Protocol } from './protocol/Commands';

export class Connection {
  private readonly server: string;
  private readonly port: number;
  private readonly queue: PacketQueue = new PacketQueue();
  private readonly errorHandler: (error: Error) => void;
  private socket: net.Socket;

  constructor(server: string, port: number, errorHandler: (error: Error) => void) {
    this.server = server;
    this.port = port;
    this.errorHandler = errorHandler;
    this.socket = new net.Socket();
    this.socket.on('error', this.errorHandler);
    this.socket.connect(this.port, this.server);
    this.socket.addListener('data', this.receivedData.bind(this));
  }

  private receivedData(data: Buffer) {
    try {
      const raw = new TextDecoder('utf-8').decode(data);

      const packetUUID = new PacketDeserializer(raw).scanUUID();
      const wrappedPacket = this.queue.popPacket(packetUUID);
      if (wrappedPacket) {
        Connection.handleAnswer(wrappedPacket, raw);
      }
    } catch (e) {
      this.errorHandler(e);
    }
  }

  private static handleAnswer(wrappedQuestionPacket: EnqueuedPacket, rawAnswer: string) {
    try {
      const command = wrappedQuestionPacket.packet.message.instruction as ClientProtocolInstruction<ProtoPayload>;
      const payloadClass = command.expectedResponsePayloadClass;
      const answer = new PacketDeserializer(rawAnswer, payloadClass).deserialize();
      if (answer.message.instruction.domain === Protocol.PossibleServerInstructions.FAILURE.domain) {
        wrappedQuestionPacket.reject(answer);
      } else {
        wrappedQuestionPacket.resolve(answer.message);
      }
    } catch (e) {
      wrappedQuestionPacket.reject(e);
    }
  }

  send<T extends ProtoPayload>(message: Message<ProtoPayload>): Promise<Message<T>> {
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
