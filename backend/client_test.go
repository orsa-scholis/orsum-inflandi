package main

import (
	server2 "github.com/orsa-scholis/orsum-inflandi-II/backend/server"
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	server, _ := server2.initServer(false)
	socket, _ := net.Listen("tcp", ":4560")

	conn, _ := socket.Accept()
	go makeConnection("Lukas")

	client := server2.initClient("Lukas", conn, server)

	if client.name != "Lukas" {
		t.Errorf("User.name = %s; want 'Lukas'", client.name)
	}
}

func makeConnection(name string) {
	net.Dial("tcp", ":4560")
}
