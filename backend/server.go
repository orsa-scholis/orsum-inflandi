package main

import (
	"fmt"
	"github.com/google/logger"
	"net"
)

type ServerMessage struct {
	domain      string
	command     string
	param       string
	attachement string
}

type Server struct {
	verbose     bool
	clients     []Client
	stoppedChan chan bool
	socket      net.Listener
}

func InitServer(verbose bool) (server Server, err error) {
	socket, err := net.Listen("tcp", ":4560")
	if nil != err {
		return
	}

	server = Server{
		verbose:     verbose,
		clients:     make([]Client, 0),
		stoppedChan: make(chan bool, 1),
		socket:      socket,
	}

	return
}

func (s *Server) Start() {
	logger.Info("Server started listening")

	for {
		conn, err := s.socket.Accept()

		if nil != err {
			select {
			case <-s.stoppedChan:
				logger.Info("Server is stopped, can't accept anymore connections")
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
	newClient := Client{
		Name: fmt.Sprintf("Client Nr. %v", len(s.clients)),
		Conn: connection,
	}
	s.clients = append(s.clients, newClient)
	newClient.handleConnection()
}

func (s *Server) CleanUp() error {
	s.stoppedChan <- true
	logger.Infof("Sending closing calls to %v clients\n", len(s.clients))

	for i, client := range s.clients {
		_, err := client.Conn.Write([]byte("error:closed\n"))
		if err != nil {
			return err
		}
		err = client.Conn.Close()
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
