import ProtocolInstruction from './ProtocolInstruction';
import ProtobufMessageClass from '../ProtobufMessageClass';

export default class ClientProtocolInstruction<T> extends ProtocolInstruction {
  readonly expectedResponsePayloadClass?: ProtobufMessageClass<T>;

  constructor(domain: string, command?: string, expectedResponsePayloadClass?: ProtobufMessageClass<T>) {
    super(domain, command);
    this.expectedResponsePayloadClass = expectedResponsePayloadClass;
  }
}
