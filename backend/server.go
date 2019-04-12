package main

import (
	"fmt"
	"github.com/google/logger"
	"net"
)

//type ServerMessage struct {
//	domain      string
//	command     string
//	param       string
//	attachment string
//}

type server struct {
	verbose     bool
	clients     []client
	games       []game
	stoppedChan chan bool
	socket      net.Listener
}

func initServer(verbose bool) (ser *server, err error) {
	socket, err := net.Listen("tcp", ":4560")
	if nil != err {
		return
	}

	ser = &server{
		verbose:     verbose,
		clients:     make([]client, 0),
		stoppedChan: make(chan bool, 1),
		socket:      socket,
	}

	return
}

func (s *server) start() {
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

func (s *server) initClientConnection(connection net.Conn) {
	newClient := initClient(fmt.Sprintf("Client #%v", len(s.clients)), connection)
	s.clients = append(s.clients, *newClient)
	newClient.start()
}

func (s *server) getGamesAsString() (list string) {
	for _, g := range s.games {
		list = list + g.name + ","
	}
	return
}

func (s *server) openGame(name string, clientOne client) bool {
	for _, g := range s.games {
		if g.name == name {
			return false
		}
	}
	s.games = append(s.games, game{
		name:      name,
		playerOne: clientOne,
		turn:      1,
	})

	return true
}

func (s *server) CleanUp() error {
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
		logger.Infof("Sent closing calls and closed socket of client #%v\n", i)
	}

	err := s.socket.Close()
	if err != nil {
		return err
	}

	return nil
}
