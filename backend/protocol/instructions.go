package protocol

import "github.com/orsa-scholis/orsum-inflandi-II/backend/legacyServer"

type ConnectionConnect struct {
	ClientInstruction
}

func(*ConnectionConnect) Execute(server *legacyServer.Server, game legacyServer.Game) bool {
	return true
}

//var Instructions = []ClientInstruction {
//
//{
//Instruction: Instruction{
//		domain: "connection",
//		command: "connect",
//	},
//	failureInstruction: Instruction{
//		domain: "failure",
//	},
//	successInstruction: {
//		Instruction: Instruction{
//			domain: "connection",
//			command: "success"
//		},
//
//	},
//}
//
//}
