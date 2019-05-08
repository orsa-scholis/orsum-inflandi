package main

type game struct {
	playerOne    *client
	playerTwo    *client
	playerOnTurn *client
	name         string
	board        *board
}

func initGame(name string) (newGame *game) {
	newGame = &game{
		name:  name,
		board: initBoard(),
	}
	newGame.board.game = newGame
	return
}

func (g *game) isFull() bool {
	if g.playerOne == nil || g.playerTwo == nil {
		return false
	}
	return true
}

func (g *game) join(client *client) clientState {
	if g.isFull() {
		return inLobby
	}

	if g.playerOne == nil {
		g.playerOne = client
		g.playerOnTurn = client
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

func (g *game) broadcastMessage(m message, me *client) {
	if !g.isFull() {
		return
	}

	if g.playerOne == me {
		g.playerTwo.sendChan <- m
	} else {
		g.playerOne.sendChan <- m
	}
}
