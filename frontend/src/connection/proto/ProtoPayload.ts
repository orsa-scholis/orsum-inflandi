import * as Protobuf from 'google-protobuf';

export interface ProtoPayload extends Protobuf.Message {
  serializeBinary(): Uint8Array;
}
