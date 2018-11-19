package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/orsa-scholis/orsum-inflandi-II/backend/server"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanUp()
	}()

	server.Start()
}

func cleanUp() {
	err := server.CleanUp()
	if nil != err {
		os.Exit(1)
	}

	fmt.Println("Shutdown gracefuly")
	os.Exit(0)
}
