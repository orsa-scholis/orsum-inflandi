package fourinarow

import (
	"github.com/orsa-scholis/orsum-inflandi-II/backend/legacyServer"
	Types "github.com/orsa-scholis/orsum-inflandi-II/proto"
)

type FourInARow struct {
	Types.Game
	Initiator    *legacyServer.Client
	Opponent     *legacyServer.Client
	playerOnTurn *legacyServer.Client
	board        *board
}

func NewGame(name string) (newFourInARow *FourInARow) {
	newGame := Types.Game{
		Name: name,
	}

	newFourInARow = &FourInARow{
		Game:  newGame,
		board: initBoard(),
	}

	newFourInARow.board.game = newFourInARow
	return
}

func (g *FourInARow) isFull() bool {
	if g.Initiator == nil || g.Opponent == nil {
		return false
	}
	return true
}

func (g *FourInARow) Join(client *legacyServer.Client) legacyServer.ClientState {
	if g.isFull() {
		return legacyServer.InLobby
	}

	if g.Initiator == nil {
		g.Initiator = client
		g.playerOnTurn = client
		return legacyServer.InGame
	}
	g.Opponent = client
	return legacyServer.PlayingGame
}

func (g *FourInARow) setStone(c *legacyServer.Client, rowNr int) bool {
	if g.board.rowFull(rowNr) {
		return false
	}

	return true
}

func (g *FourInARow) broadcastMessage(m legacyServer.UserMessage, me *legacyServer.Client) {
	if !g.isFull() {
		return
	}

	if g.Initiator == me {
		g.Opponent.SendMessage(m)
	} else {
		g.Initiator.SendMessage(m)
	}
}
