import { Message as ProtobufMessage } from 'google-protobuf';

export default interface ProtobufMessageClass<T> {
  new(...args: any[]): T;
  deserializeBinary(bytes: Uint8Array): ProtobufMessage;
}
