package server

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"encoding/asn1"
	b64 "encoding/base64"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

type Client struct {
	Name string
	Conn net.Conn
}

type ServerMessage struct {
	domain      string
	command     string
	param       string
	attachement string
}

var clients []Client
var ln net.Listener
var legacy bool

func Start() {
	flag.BoolVar(&legacy, "legacy", false, "should run in legacy mode?")

	flag.Parse()

	clients := make([]Client, 0)
	fmt.Println("Launching server...")

	var err error
	ln, err = net.Listen("tcp", ":4560")
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
			newClient := Client{
				Name: fmt.Sprintf("Client Nr. %v", len(clients)),
				Conn: conn,
			}
			clients = append(clients, newClient)
			go handleConnection(newClient)
		}
	}
}

func handleConnection(client Client) {
	reader := rand.Reader
	bitSize := 2048

	key, keyErr := rsa.GenerateKey(reader, bitSize)
	publicKey := key.PublicKey

	if keyErr != nil {
		fmt.Println(keyErr)
	}

	var conn = client.Conn

	message, _ := bufio.NewReader(client.Conn).ReadString('\n')

	// TODO: Implement better check func
	messageParts := strings.Split(message, ":")
	if len(messageParts) != 3 {
		fmt.Println("First message too short")
	}
	if (messageParts[0]) != "connection" {
		fmt.Println("First message domain wrong")
	}
	if (messageParts[1]) != "connect" {
		fmt.Println("First message action wrong")
	}
	if (len(messageParts[2])) < 4 {
		fmt.Println("First message param too short")
	}

	fmt.Print("Message Received:", string(message))

	if legacy {
		message = fmt.Sprintf("success:accepted:rO0ABXNyABRqYXZhLnNlY3VyaXR5LktleVJlcL35T7OImqVDAgAETAAJYWxnb3JpdGhtdAASTGphdmEvbGFuZy9TdHJpbmc7WwAHZW5jb2RlZHQAAltCTAAGZm9ybWF0cQB+AAFMAAR0eXBldAAbTGphdmEvc2VjdXJpdHkvS2V5UmVwJFR5cGU7eHB0AANSU0F1cgACW0Ks8xf4BghU4AIAAHhwAAAAojCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAogfCMH/o4tfQHdN6ABXw8w02kbUJE4XQSfYq43kRu9BbRZOF4bCCBrL4G4lqMGmcuJiP5wSIpkjFG5C+BnZx1OLghC/kOThyhhI+cfDL2hgr/It8ELrdw3CqhugGfRzZJQD9ZkjdvceD2Wts00cyPfpJTAG+KVFeEJcg1knFN5ECAwEAAXQABVguNTA5fnIAGWphdmEuc2VjdXJpdHkuS2V5UmVwJFR5cGUAAAAAAAAAABIAAHhyAA5qYXZhLmxhbmcuRW51bQAAAAAAAAAAEgAAeHB0AAZQVUJMSUM=")
	} else {
		// TODO: Implement correct key exchange
		publicBytes, err := asn1.Marshal(publicKey)
		if nil != err {
			panic(err)
		}
		var publicAs64 = b64.StdEncoding.EncodeToString(publicBytes)
		message = fmt.Sprintf("success:accepted:%v\n", publicAs64)
	}
	_, err := conn.Write([]byte(message))
	if nil != err {
		fmt.Println("No writy writy")
	} else {
		fmt.Printf("Message Sent: '%v'\n", message)
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

	err := ln.Close()
	if nil != err {
		return err
	}

	return nil
}
