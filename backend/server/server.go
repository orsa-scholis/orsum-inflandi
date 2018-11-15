package server

import (
  "bufio"
  "fmt"
  "net"
  "os"
  "strings"
)

func StartServer() {
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
      go handleConnection(conn)
    }
  }
  fmt.Println("terminated")
}

func handleConnection(conn net.Conn) {
  message, _ := bufio.NewReader(conn).ReadString('\n')
  fmt.Print("Message Received:", string(message))
  newmessage := strings.ToUpper(message)
  conn.Write([]byte(newmessage + "\n"))
}
