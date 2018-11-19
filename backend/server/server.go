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

	defer ln.Close()

	for {
		conn, err := ln.Accept()

		if nil != err {
			fmt.Fprint(os.Stderr, "can't accept connection: ", err)
		} else {
			var newClient Client
			newClient.Name = fmt.Sprintf("Client Nr. %v", len(clients))
			newClient.Conn = conn
			clients = append(clients, newClient)
			go handleConnection(newClient)
		}
	}
}

func handleConnection(client Client) {
	var conn = client.Conn

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message Received:", string(message))
	newmessage := strings.ToUpper(message)
	_, err := conn.Write([]byte(newmessage + "\n"))
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
		err2 := conn.Close()
		if nil != err2 {
			return err2
		}
	}

	return nil
}
