package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Hardcoding config parameters, however these should
// be handled via a config file in a prod system.
const (
	grpcServerAddr    = "127.0.0.1:30000"
	grpcServerUseTLS  = true
	grpcServerCrtFile = "worker.crt"
	grpcServerKeyFile = "worker.key"

	schedulerAddr    = "127.0.0.1:50000"
	schedulerCrtFile = "scheduler.crt"
)

var (
	// workerID is the id assigned by the scheduler
	// after registering on scheduler.
	workerID string

	// apiKey is used for scheduler and worker communications.
	apiKey string

	errs = make(chan error)
	sig  = make(chan os.Signal)
)

func init() {
	flag.StringVar(&apiKey, "api_key", "", "API key for the worker and scheduler api communication.")
	flag.Parse()

	if apiKey == "" {
		log.Fatalln("API key was not specified")
	}

	checkKeyCrt()
}

// Entry point of the worker application.
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go startGRPCServer()
	go registerWorker(ctx)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errs:
		deregisterWorker(ctx)
		log.Fatalln(err)
	case s := <-sig:
		deregisterWorker(ctx)
		log.Printf("Signal (%d) received, stopping\n", s)
		os.Exit(0)
	}
}
