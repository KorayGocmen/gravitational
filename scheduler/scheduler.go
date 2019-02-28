package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Hardcoding config parameters, however these should
// be handled via a config file in a prod system.
const (
	grpcServerAddr    = "127.0.0.1:50000"
	grpcServerUseTLS  = true
	grpcServerCrtFile = "scheduler.crt"
	grpcServerKeyFile = "scheduler.key"
)

var (
	errs = make(chan error)
	sig  = make(chan os.Signal)

	apiKey string
)

func init() {
	flag.StringVar(&apiKey, "api_key", "", "API key for the worker and scheduler api communication.")
	flag.Parse()

	if apiKey == "" {
		errs <- errors.New("api key was not set")
	}

	checkKeyCrt()
}

// Entry point of the scheduler application.
func main() {

	go startGRPCServer()

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errs:
		log.Fatalln(err)
	case s := <-sig:
		log.Printf("Signal (%d) received, stopping\n", s)
		os.Exit(0)
	}
}
