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
    fmt.Fprintf(os.Stderr, "Cannot shut down cleanly: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Shutdown gracefully")
	os.Exit(0)
}
