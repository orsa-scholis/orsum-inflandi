package main

import (
	"flag"
	"fmt"
	"github.com/google/logger"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

func main() {
	flag.Parse()

	defer logger.Init("OrsumInflandiII-Backend", *verbose, false, ioutil.Discard).Close()

	server, err := initServer(*verbose)

	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		server.start()
	}()

	logger.Info("Starting server...")

	<-sigChan
	cleanUp(*server)
}

func cleanUp(server server) {
	err := server.CleanUp()
	if nil != err {
		fmt.Fprintf(os.Stderr, "Cannot shut down cleanly: %v\n", err)
		os.Exit(1)
	}

	logger.Info("Shutdown gracefully")
	os.Exit(0)
}
