package main

type game struct {
	playerOne *client
	playerTwo *client
	turn      int
	name      string
	board     *board
}

func initGame(name string, playerOne *client) (newGame *game) {
	newGame = &game{
		name:      name,
		playerOne: playerOne,
		turn:      1,
		board:     initBoard(),
	}
	newGame.board.game = newGame
	return
}

func (g *game) isFull() bool {
	if g.playerOne.name == "" || g.playerTwo.name == "" {
		return false
	}
	return true
}

func (g *game) join(client *client) clientState {
	if g.isFull() {
		return inLobby
	}

	if g.playerOne.name == "" {
		g.playerOne = client
		return inGame
	}
	g.playerTwo = client
	return playingGame
}

func (g *game) setStone(c *client, rowNr int) bool {
	if g.board.rowFull(rowNr) {
		return false
	}

	return true
}
