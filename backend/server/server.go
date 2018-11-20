package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Client struct {
	Name string
	Conn net.Conn
}

var clients []Client

func Start() {
	clients := make([]Client, 0)
	fmt.Println("Launching server...")

	ln, err := net.Listen("tcp", ":4560")
	if nil != err {
		fmt.Fprint(os.Stderr, "can't create TCP socket: ", err)
		os.Exit(1)
	}

	fmt.Println("Server started listening")

	for {
		conn, err := ln.Accept()

		if nil != err {
			fmt.Fprint(os.Stderr, "can't accept connection: ", err)
		} else {
			newClient := Client {
			  Name: fmt.Sprintf("Client Nr. %v", len(clients)),
			  Conn: conn,
      }
			clients = append(clients, newClient)
			go handleConnection(newClient)
		}
	}
}

func handleConnection(client Client) {
  message, _ := bufio.NewReader(client.Conn).ReadString('\n')
	fmt.Print("Message Received:", string(message))

	newMessage := strings.ToUpper(message)
	_, err := client.Conn.Write([]byte(newMessage + "\n"))

	if nil != err {
		fmt.Println("No writy writy")
	}
}

func CleanUp() error {
	for clientI := 0; clientI < len(clients); clientI++ {
		var conn = clients[clientI].Conn

		_, err := conn.Write([]byte("error:closed\n"))
		if nil != err {
			return err
		}

		err = conn.Close()
		if nil != err {
			return err
		}
	}

	return nil
}
