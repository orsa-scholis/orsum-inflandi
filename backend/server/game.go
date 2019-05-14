package server

import "github.com/google/uuid"

type GameType int

const (
	fourInARow GameType = iota
)

func (gType GameType) String() string {
	gTypes := [...]string{
		"4 in a row",
	}

	if gType < 1 || gType > 1 {
		return "Unknown GameType"
	}

	return gTypes[gType]
}

type Game interface {
	GetUUID() uuid.UUID
	GetName() string
	isFull() bool
	Join(user *User) UserState
	broadcastMessage(m UserMessage, me *User) // TODO: move to server
	GetGameType() GameType
}
