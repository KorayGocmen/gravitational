package main

import (
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

	defAPIKey = "default-api-key"
)

var (
	errs = make(chan error)
	sig  = make(chan os.Signal)

	apiKey string
)

func init() {
	flag.StringVar(&apiKey, "api_key", defAPIKey, "API key for the worker and scheduler api communication.")
	flag.Parse()

	checkKeyCrt()
}

// Entry point of the scheduler application.
func main() {

	go startGRPCServer(server{})

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errs:
		log.Fatalln(err)
	case s := <-sig:
		log.Printf("Signal (%d) received, stopping\n", s)
		os.Exit(0)
	}
}
