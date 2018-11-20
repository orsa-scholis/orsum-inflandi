package main

import (
	"bufio"
	"fmt"
	"github.com/google/logger"
	"net"
	"strings"
)

type client struct {
	Name string
	Conn net.Conn
}

func (c *client) receiveMessage() (message string) {
	message, _ = bufio.NewReader(c.Conn).ReadString('\n')
	return
}

func (c *client) sendMessage(message string) error {
	_, err := c.Conn.Write([]byte(message))
	return err
}

func (c *client) handleConnection() {
	message := c.receiveMessage()

	return
}

func validate(received message, expected message) bool {
	if !received.receive {
		return false
	}
	if (messageParts[0]) != "connection" {
		logger.Warning("First message domain wrong")
	}
	if (messageParts[1]) != "connect" {
		logger.Warning("First message action wrong")
	}
	if (len(messageParts[2])) < 4 {
		logger.Warning("First message param too short")
	}

	logger.Info("Message Received:", string(message))

	err := c.sendMessage("success:accepted\n")
	if err != nil {
		fmt.Println("No writy writy")
	} else {
		fmt.Printf("Message Sent: '%v'\n", message)
	}
}

func (c *client) receiver() {
	for {
		raw, _, err := bufio.NewReader(c.conn).ReadLine()

		if err != nil {
			select {
			case <-c.stoppedChan:
				logger.Info("server is stopped, can't accept anymore messages")
				return
			default:
				logger.Warning("can't accept message: ", err)
				return
			}
		}

		parsed, err := parseMessage(string(raw))
		if err != nil {
			c.sendChan <- message{
				receive: false,
				domain:  "error",
				command: "message",
				param:   err.Error(),
			}
		}
		logger.Info(fmt.Sprintf("Client '%s' received message: '%+v'", c.name, parsed))
		c.receiveChan <- parsed
	}
}

func (c *client) sender() {
	for {
		sendMessage := <-c.sendChan
		_, err := c.conn.Write([]byte(sendMessage.String() + "\n"))

		if err != nil {
			select {
			case <-c.stoppedChan:
				logger.Info("server is stopped, can't accept anymore messages")
				return
			default:
				logger.Warning("can't write message: ", err)
				continue
			}
		}

		logger.Infof(fmt.Sprintf("Client '%s' sent message: '%+v'", c.name, sendMessage))
	}
}

func (c *client) start() {
	for {
		cMessage := <-c.receiveChan
		messageHandled := false

		if c.state == connecting && validate(cMessage, connectionConnect) {
			c.name = cMessage.param
			c.state = inLobby

			c.sendChan <- message{
				receive: false,
				domain:  "success",
				command: "accepted",
				param:   "",
			}

			messageHandled = true
		}

		if c.state == inLobby && cMessage.matches(infoRequestGames) {
			games := c.server.getGamesAsString()

			c.sendChan <- message{
				receive: false,
				domain:  "success",
				command: "requested",
				param:   games,
			}

			messageHandled = true
		}

		if c.state == inLobby && cMessage.matches(serverNewGame) {
			success := c.server.openGame(cMessage.param, *c)

			if success {
				c.sendChan <- message{
					receive: false,
					domain:  "success",
					command: "created",
				}
			} else {
				c.sendChan <- message{
					receive: false,
					domain:  "error",
					command: "newGame",
					param:   "game name not unique",
				}
			}

			messageHandled = true
		}


		if !messageHandled {
			logger.Warning("unhandled message: '%s'", cMessage)
		}
	}
}

var connectionConnect = message{
	receive: true,
	domain:  "connection",
	command: "connect",
	param:   "*",
}

var infoRequestGames = message{
	receive: true,
	domain:  "info",
	command: "requestGames",
	param:   "*",
}

var serverNewGame = message{
	receive: true,
	domain:  "server",
	command: "newGame",
	param:   "*|required,min:3,max:5",
}
