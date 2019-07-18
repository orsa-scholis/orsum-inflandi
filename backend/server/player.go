package server

import (
	Types "github.com/orsa-scholis/orsum-inflandi-II/proto"
	"net"
)

type playerMessage struct {
	id      uuid.UUID
	domain  string
	command string
	param   string
}

type player struct {
	Types.User
	conn        net.Conn
	stoppedChan chan bool
	receiveChan chan playerMessage
	sendChan    chan serverMessage
	state       ClientState
	server      *Server
	currentGame uuid.UUID
}

func newPlayer(name string, conn net.Conn, server *Server) (pl *player) {
	user = &Types.User{ugins
		name,
	}

	pl = &player{

		conn:        conn,
		stoppedChan: make(chan bool, 1),
		receiveChan: make(chan UserMessage),
		sendChan:    make(chan UserMessage, 10),
		state:       Connecting,
		server:      server,
	}

	go c.receiver()
	go c.sender()

	return
}
