package main

import (
	"bufio"
	"fmt"
	"github.com/google/logger"
	"net"
	"strings"
)

type Client struct {
	Name string
	Conn net.Conn
}

func (c *Client) receiveMessage() (message string) {
	message, _ = bufio.NewReader(c.Conn).ReadString('\n')
	return
}

func (c *Client) sendMessage(message string) error {
	_, err := c.Conn.Write([]byte(message))
	return err
}

func (c *Client) handleConnection() {
	message := c.receiveMessage()

	// TODO: Implement better check func
	messageParts := strings.Split(message, ":")
	if len(messageParts) != 3 {
		logger.Warning("First message too short")
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
