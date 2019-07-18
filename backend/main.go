package main

import (
	"flag"
	"fmt"
	"github.com/google/logger"
	"github.com/orsa-scholis/orsum-inflandi-II/backend/server"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

func main() {
	flag.Parse()

	defer logger.Init("OrsumInflandiII-Backend", *verbose, false, ioutil.Discard).Close()

	gameServer, err := server.InitServer(*verbose)

	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		gameServer.Start()
	}()

	logger.Info("Starting backend...")

	<-sigChan
	cleanUp(*gameServer)
}

func cleanUp(server server.Server) {
	//err := server.CleanUp()
	//if nil != err {
	//	_, _ = fmt.Fprintf(os.Stderr, "Cannot shut down cleanly: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//logger.Info("Shutdown gracefully")
	//os.Exit(0)
}
