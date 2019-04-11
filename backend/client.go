package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/google/logger"
	"net"
	"strings"
)

type clientState int

const (
	connecting clientState = iota
	inLobby
	inGame
	playingGame
	disconnected
)

func (state clientState) String() string {
	states := [...]string{
		"Connecting",
		"InLobby",
		"InGame",
		"Disconnected"}

	if state < connecting || state > disconnected {
		return "Unknown ClientState"
	}

	return states[state]
}

type message struct {
	receive    bool
	domain     string
	command    string
	param      string
	attachment string
}

func (m *message) String() string {
	return fmt.Sprintf("%s:%s:%s", m.domain, m.command, m.param)
}

func parseMessage(raw string) (m message, err error) {
	messageParts := strings.Split(raw, ":")
	if len(messageParts) == 3 {
		m = message{
			receive: true,
			domain:  messageParts[0],
			command: messageParts[1],
			param:   messageParts[2],
		}
	}
	if len(messageParts) == 2 {
		m = message{
			receive: true,
			domain:  messageParts[0],
			command: messageParts[1],
		}
	}
	if len(messageParts) < 2 || len(messageParts) > 3 {
		err = errors.New("message length invalid")
	}

	return
}

func (received *message) matches(expected message) bool {
	if !received.receive {
		return false
	}
	if expected.domain == "*" {
		return true
	}
	if received.domain != expected.domain {
		return false
	}
	if received.command != expected.command {
		return false
	}
	if received.param != expected.param && expected.param != "*" {
		return false
	}
	if expected.param[0] == '*' && len(expected.param) > 1 {
		rules := expected.param[2:]
		ruleList := strings.Split(rules, ",")
		if len(ruleList) == 0 {
			return true
		}
		for _, rule := range(ruleList) {
			ruleParts := strings.Split(rule, ":")
			if len(ruleParts) > 1 {
				if !validatorsWithParam[rule](received.param, ruleParts[1]) {
					return false
				}
			} else {
				if !validators[rule](received.param) {
					return false
				}
			}
		}

	}
	return true
}

type client struct {
	name        string
	conn        net.Conn
	stoppedChan chan bool
	receiveChan chan message
	sendChan    chan message
	state       clientState
	server      server
}

func initClient(name string, conn net.Conn) (c *client) {
	c = &client{
		name:        name,
		conn:        conn,
		stoppedChan: make(chan bool, 1),
		receiveChan: make(chan message),
		sendChan:    make(chan message, 10),
		state:       connecting,
	}

	go c.receiver()
	go c.sender()

	return
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

		if c.state == connecting && cMessage.matches(connectionConnect) {
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
