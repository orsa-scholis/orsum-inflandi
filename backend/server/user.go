package server

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/google/logger"
	"github.com/google/uuid"
	"net"
	"strconv"
	"strings"
)

type UserState int

const (
	Connecting UserState = iota
	InLobby
	InGame
	PlayingGame
	Disconnected
)

func (state UserState) String() string {
	states := [...]string{
		"Connecting",
		"InLobby",
		"InGame",
		"Disconnected"}

	if state < Connecting || state > Disconnected {
		return "Unknown ClientState"
	}

	return states[state]
}

type UserMessage struct {
	id      uuid.UUID
	domain  string
	command string
	param   string
}

func (m *UserMessage) String() string {
	return fmt.Sprintf("%s:%s:%s", m.domain, m.command, m.param)
}

func parseMessage(raw string) (m UserMessage, err error) {
	messageParts := strings.Split(raw, ":")
	if len(messageParts) == 3 {
		m = UserMessage{
			domain:  messageParts[0],
			command: messageParts[1],
			param:   messageParts[2],
		}
	}
	if len(messageParts) == 2 {
		m = UserMessage{
			domain:  messageParts[0],
			command: messageParts[1],
			param:   "",
		}
	}
	if len(messageParts) < 2 || len(messageParts) > 3 {
		err = errors.New("UserMessage length invalid")
	}

	return
}

type messageValidationError struct {
	err            string
	returnToClient bool
}

func (messageErr *messageValidationError) Error() string {
	return messageErr.err
}

func validateMessageDomainCommand(received UserMessage, expected UserMessage) (errs []error) {
	if expected.domain == "*" {
		return
	}

	if received.domain != expected.domain {
		errs = append(errs, &messageValidationError{err: fmt.Sprintf("the domain '%v' does not match expected '%v'", received.domain, expected.domain)})
	}
	if received.command != expected.command {
		errs = append(errs, &messageValidationError{err: fmt.Sprintf("the command '%v' does not match expected '%v'", received.command, expected.command)})
	}
	return errs
}

func validateMessageParam(received UserMessage, expected UserMessage) error {
	if expected.param[0] == '*' && len(expected.param) > 1 {
		rules := expected.param[2:]
		ruleList := strings.Split(rules, ",")
		if len(ruleList) > 0 {
			ruleErrors := ""
			for _, rule := range ruleList {
				ruleParts := strings.Split(rule, ";")

				var ruleError error
				if len(ruleParts) > 1 {
					logger.Infof("executing param checker '%s' with attr '%v' and param '%s'", ruleParts[0], ruleParts[1], received.param)
					ruleError = validatorsWithParam[ruleParts[0]](received.param, ruleParts[1])
				} else {
					logger.Infof("executing param checker '%s' with param '%s'", ruleParts[0], received.param)
					ruleError = validators[ruleParts[0]](received.param)
				}

				if ruleError != nil {
					logger.Infof("the param '%s' does not match rule '%s'", received.param, rule)
					ruleErrors += rule + ","
				}
			}

			if ruleErrors != "" {
				fullErrorMessage := fmt.Sprintf("%v", ruleErrors)
				return &messageValidationError{err: fullErrorMessage[:len(fullErrorMessage)-1], returnToClient: true}
			}
		}
	}

	return nil
}

type User struct {
	name        string
	conn        net.Conn
	stoppedChan chan bool
	receiveChan chan UserMessage
	sendChan    chan UserMessage
	state       UserState
	server      *Server
	currentGame uuid.UUID
}

func newUser(name string, conn net.Conn, server *Server) (c *User) {
	c = &User{
		name:        name,
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

func (c *User) receiver() {
	for {
		raw, _, err := bufio.NewReader(c.conn).ReadLine()

		if err != nil {
			select {
			case <-c.stoppedChan:
				logger.Info("server is stopped, can't accept anymore messages")
				return
			default:
				logger.Warning("can't accept UserMessage: ", err)
				return
			}
		}

		parsed, err := parseMessage(string(raw))
		if err != nil {
			c.sendChan <- UserMessage{
				domain:  "error",
				command: "UserMessage",
				param:   err.Error(),
			}
		}
		logger.Info(fmt.Sprintf("Client '%s' received UserMessage: '%+v'", c.name, parsed))
		c.receiveChan <- parsed
	}
}

func (c *User) sender() {
	for {
		sendMessage := <-c.sendChan
		_, err := c.conn.Write([]byte(sendMessage.String() + "\n"))

		if err != nil {
			select {
			case <-c.stoppedChan:
				logger.Info("server is stopped, can't accept anymore messages")
				return
			default:
				logger.Warning("can't write UserMessage: ", err)
				continue
			}
		}

		logger.Infof(fmt.Sprintf("Client '%s' sent UserMessage: '%+v'", c.name, sendMessage))
	}
}

func (c *User) start() {
	for {
		cMessage := <-c.receiveChan

		messageHandled, messageError := c.checkMessageHandlersIfFit(cMessage)

		if !messageHandled {
			logger.Warningf("unhandled UserMessage: '%s', \nerrors: '%v+'\nstate: '%s'", cMessage.String(), messageError, c.state)
			c.sendChan <- UserMessage{
				domain:  "error",
				command: cMessage.command,
				param:   messageError,
			}
		}
	}
}

func (c *User) checkMessageHandlersIfFit(cMessage UserMessage) (messageHandled bool, messageError string) {
	stateFunctions := clientStateMessageHandlers[c.state]
	for expMessage, messageFunc := range stateFunctions {
		errs := validateMessageDomainCommand(cMessage, expMessage)
		if len(errs) > 0 {
			logger.Infof("%+v", errs)
			continue
		}
		if (cMessage.param == "" || len(cMessage.param) == 0) && expMessage.param == "*" {
			continue
		}

		paramErr := validateMessageParam(cMessage, expMessage)
		errs = append(errs, paramErr)
		if paramErr == nil {
			messageFunc(c, cMessage)
			logger.Infof("Client '%s' handled UserMessage: '%+v', with messageHandler: '%s'", c.name, cMessage, expMessage.String())

			messageHandled = true
			break
		} else {
			if err, ok := paramErr.(*messageValidationError); ok && err.returnToClient {
				messageError = err.Error()
			}
		}
	}
	return
}

func (c *User) SendMessage(message UserMessage) {
	c.sendChan <- message
}

func connectionConnectHandler(c *User, recMessage UserMessage) {
	c.name = recMessage.param
	c.state = InLobby

	c.sendChan <- UserMessage{
		domain:  "success",
		command: "accepted",
		param:   "",
	}
}

func infoRequestGamesHandler(c *User, _ UserMessage) {
	games := c.server.getGamesAsString()

	c.sendChan <- UserMessage{
		domain:  "success",
		command: "requested",
		param:   games,
	}
}

func serverNewGameHandler(c *User, recMessage UserMessage) {
	success := c.server.openGame(recMessage.param, c)

	if success {
		c.sendChan <- UserMessage{
			domain:  "success",
			command: "created",
			param:   recMessage.param,
		}

		// TODO: Change to full gameList with playerCount
		go c.server.broadcastMessage(UserMessage{
			domain:  "subscription",
			command: "gameAdded",
			param:   c.server.getGamesAsString(),
		}, c)
	} else {
		c.sendChan <- UserMessage{
			domain:  "error",
			command: "newGame",
			param:   "game name not unique",
		}
	}
}

func gameJoinHandler(c *User, recMessage UserMessage) {
	gameID, _ := strconv.Atoi(recMessage.param)
	clientState := c.server.joinGame(gameID, c)

	if clientState == InLobby {
		c.sendChan <- UserMessage{
			domain:  "error",
			command: "join",
			param:   "game full or non existent",
		}

		return
	}

	c.state = clientState
	c.currentGame = gameID

	var successMessageTemplate = UserMessage{
		domain:  "success",
		command: "joined",
	}

	if clientState == InGame {
		successMessageTemplate.param = "1"
		c.sendChan <- successMessageTemplate
	} else if clientState == PlayingGame {
		successMessageTemplate.param = "2"
		c.sendChan <- successMessageTemplate
	}
}

func gameSetStoneHandler(c *User, recMessage UserMessage) {
	rowNr, _ := strconv.Atoi(recMessage.param)
	success := c.server.getGame(c.currentGame).setStone(c, rowNr)

	if success {
		c.sendChan <- UserMessage{
			domain:  "success",
			command: "setStone",
			param:   "",
		}
	} else {
		c.sendChan <- UserMessage{
			domain:  "error",
			command: "setStone",
			param:   "",
		}
	}
}

var clientStateMessageHandlers = map[UserState]map[UserMessage]func(*User, UserMessage){
	Connecting: {
		connectionConnect: connectionConnectHandler,
	},
	InLobby: {
		infoRequestGames: infoRequestGamesHandler,
		serverNewGame:    serverNewGameHandler,
		gameJoin:         gameJoinHandler,
	},
	InGame: {
		gameSetStone: gameSetStoneHandler,
	},
	PlayingGame: {},
}

var connectionConnect = UserMessage{
	domain:  "connection",
	command: "connect",
	param:   "*",
}

var infoRequestGames = UserMessage{
	domain:  "info",
	command: "requestGames",
	param:   "*",
}

var serverNewGame = UserMessage{
	domain:  "server",
	command: "newGame",
	param:   "*|required,gt;2,lt;6",
}

var gameJoin = UserMessage{
	domain:  "game",
	command: "join",
	param:   "*|required,int",
}

var gameSetStone = UserMessage{
	domain:  "game",
	command: "setStone",
	param:   "*|required,int,gt:0,lt:8",
}
