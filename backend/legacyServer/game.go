package legacyServer

import (
	"github.com/google/uuid"
	Types "github.com/orsa-scholis/orsum-inflandi-II/proto"
)

type Game interface {
	GetUUID() uuid.UUID
	GetName() string
	isFull() bool
	Join(user *Client) ClientState
	broadcastMessage(m UserMessage, me *Client) // TODO: move to legacyServer
	GetGameType() Types.GameType
	GetGameResult() Types.GameResult
}
