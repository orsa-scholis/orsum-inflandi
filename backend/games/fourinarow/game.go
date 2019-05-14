package fourinarow

import "github.com/orsa-scholis/orsum-inflandi-II/backend/server"

type Game struct {
	playerOne    *server.User
	playerTwo    *server.User
	playerOnTurn *server.User
	Name         string
	board        *board
}

func NewGame(name string) (newGame *Game) {
	newGame = &Game{
		Name:  name,
		board: initBoard(),
	}
	newGame.board.game = newGame
	return
}

func (g *Game) isFull() bool {
	if g.playerOne == nil || g.playerTwo == nil {
		return false
	}
	return true
}

func (g *Game) Join(client *server.User) server.UserState {
	if g.isFull() {
		return server.InLobby
	}

	if g.playerOne == nil {
		g.playerOne = client
		g.playerOnTurn = client
		return server.InGame
	}
	g.playerTwo = client
	return server.PlayingGame
}

func (g *Game) setStone(c *server.User, rowNr int) bool {
	if g.board.rowFull(rowNr) {
		return false
	}

	return true
}

func (g *Game) broadcastMessage(m server.UserMessage, me *server.User) {
	if !g.isFull() {
		return
	}

	if g.playerOne == me {
		g.playerTwo.SendMessage(m)
	} else {
		g.playerOne.SendMessage(m)
	}
}
