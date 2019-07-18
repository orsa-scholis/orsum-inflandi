package protocol

import (
	//"github.com/golang/protobuf/proto"
	"github.com/orsa-scholis/orsum-inflandi-II/backend/legacyServer"
)

type ProtobufMessagePayload struct {}

type ClientInstruction interface {
	Instruction()
	successInstruction() ClientInstruction
	failureInstruction() Instruction
	Execute(legacyServer.Server, legacyServer.Game)
}

func ExecuteInstruction(clientInstruction *ClientInstruction) {
	clientInstruction.Execute(legacyServer.Server{}, legacyServer.Game())
}
