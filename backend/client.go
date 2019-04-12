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

func validateMessage(received message, expected message) (err error) {
	if !received.receive {
		err = errors.New("the message needs to be received")
	}
	if expected.domain == "*" {
		return
	}
	if received.domain != expected.domain {
		err = fmt.Errorf("the domain '%v' does not match expected '%v'", received.domain, expected.domain)
	}
	if received.command != expected.command {
		err = fmt.Errorf("the command '%v' does not match expected '%v'", received.command, expected.command)
	}
	if received.param != expected.param && expected.param != "*" {
		err = fmt.Errorf("the param '%v' does not match expected '%v'", received.param, expected.param)
	}

	if err != nil {
		return err
	}

	if expected.param[0] == '*' && len(expected.param) > 1 {
		rules := expected.param[2:]
		ruleList := strings.Split(rules, ",")
		if len(ruleList) > 0 {
			ruleErrors := ""
			for _, rule := range ruleList {
				ruleParts := strings.Split(rule, ":")

				if len(ruleParts) > 1 {
					ruleErrors += "," + validatorsWithParam[rule](received.param, ruleParts[1]).Error()
				} else {
					ruleErrors += "," + validators[rule](received.param).Error()
				}
			}
		}
	}
	return
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
		messageErrors := make([]string, 2)

		stateFunctions := clientStateMessageHandlers[c.state]
		for expMessage, messageFunc := range stateFunctions {

			if validateMessage(cMessage, expMessage) == nil {
				messageFunc(c, cMessage)

				messageHandled = true
				break
			}
		}

		if !messageHandled {
			logger.Warningf("unhandled message: '%s', \nerrors: '%v+'\nstate: '%s'", cMessage.String(), messageErrors, c.state)
		}
	}
}

func connectionConnectHandler(c *client, recMessage message) {
	c.name = recMessage.param
	c.state = inLobby

	c.sendChan <- message{
		receive: false,
		domain:  "success",
		command: "accepted",
		param:   "",
	}
}

func infoRequestGamesHandler(c *client, _ message) {
	games := c.server.getGamesAsString()

	c.sendChan <- message{
		receive: false,
		domain:  "success",
		command: "requested",
		param:   games,
	}
}

func serverNewGameHandlder(c *client, recMessage message) {
	success := c.server.openGame(recMessage.param, *c)

	if success {
		c.sendChan <- message{
			receive: false,
			domain:  "success",
			command: "created",
			param: "1",
		}
	} else {
		c.sendChan <- message{
			receive: false,
			domain:  "error",
			command: "newGame",
			param:   "game name not unique",
		}
	}
}

var clientStateMessageHandlers = map[clientState]map[message]func(*client, message){
	connecting: {
		connectionConnect: connectionConnectHandler,
	},
	inLobby: {
		infoRequestGames: infoRequestGamesHandler,
		serverNewGame: serverNewGameHandlder,
	},
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
