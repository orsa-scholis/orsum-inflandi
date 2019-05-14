package server

import (
	"fmt"
	"github.com/google/logger"
	"github.com/google/uuid"
	"net"
)

type ServerMessage struct {
	id     uuid.UUID
	status string
	param  string
}

type Server struct {
	verbose     bool
	clients     []*User
	games       map[GameType]map[uuid.UUID]Game
	stoppedChan chan bool
	socket      net.Listener
}

func InitServer(verbose bool) (ser *Server, err error) {
	socket, err := net.Listen("tcp", ":4560")
	if nil != err {
		return
	}

	ser = &Server{
		verbose:     verbose,
		clients:     make([]*User, 0),
		games:       make(map[GameType]map[uuid.UUID]Game, 0),
		stoppedChan: make(chan bool, 1),
		socket:      socket,
	}

	return
}

func (s *Server) Start() {
	logger.Info("server started listening")

	for {
		conn, err := s.socket.Accept()

		if nil != err {
			select {
			case <-s.stoppedChan:
				logger.Info("server is stopped, can't accept anymore connections")
				return
			default:
				logger.Warning("can't accept connection: ", err)
				continue
			}
		}

		go s.initClientConnection(conn)
	}
}

func (s *Server) initClientConnection(connection net.Conn) {
	newClient := newUser(fmt.Sprintf("Client #%v", len(s.clients)), connection, s)
	s.clients = append(s.clients, newClient)
	newClient.start()
}

func (s *Server) openGame(name string, clientOne *User) bool {
	for _, g := range s.games {
		if g.GetName() == name {
			return false
		}
	}

	// TODO: implement game type to newGame mapper
	//s.games = append(s.games, NewGame(name))

	return true
}

func (s *Server) joinGame(gameID uuid.UUID, user *User) (returnState UserState) {
	returnState = InLobby

	for gameType := range s.games {
		for gameUUID := range s.games[gameType] {
			if gameUUID == gameID {
				returnState = s.games[gameType][gameUUID].Join(user)
			}
		}

	}

	return
}

func (s *Server) CleanUp() error {
	s.stoppedChan <- true
	logger.Infof("Sending closing calls to %v clients\n", len(s.clients))

	for i, client := range s.clients {
		_, err := client.conn.Write([]byte("connection:closed\n"))
		if err != nil {
			return err
		}
		err = client.conn.Close()
		if err != nil {
			return err
		}
		logger.Infof("Sent closing calls and closed socket of User #%v\n", i)
	}

	err := s.socket.Close()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) broadcastMessage(m UserMessage, me *User) {
	for _, client := range s.clients {
		if client.name == me.name {
			continue
		}
		client.sendChan <- m
		logger.Infof("Sent broadcast to all clients, UserMessage: '%+v", m)
	}
}
