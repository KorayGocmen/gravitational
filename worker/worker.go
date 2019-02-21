package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Hardcoding config parameters, however these should
// be handled via a config file in a prod system.
const (
	grpcServerAddr    = "127.0.0.1:30000"
	grpcServerUseTLS  = false
	grpcServerCrtFile = "server.pem"
	grpcServerKeyFile = "server.key"

	schedulerAddr = "127.0.0.1:50000"
)

var (
	// workerID is the id assigned by the scheduler
	// after registering on scheduler.
	workerID string

	errs = make(chan error)
	sig  = make(chan os.Signal)
)

// Entry point of the worker application.
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go registerWorker(ctx)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errs:
		fatal(ctx, 1, err.Error())
	case s := <-sig:
		fatal(ctx, 0, fmt.Sprintf("Signal (%d) received, stopping\n", s))
	}
}

func fatal(ctx context.Context, code int, message string) {
	deregisterWorker(ctx)
	log.Println(message)
	os.Exit(code)
}
