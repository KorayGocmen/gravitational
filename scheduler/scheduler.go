package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Hardcoding config parameters, however these should
// be handled via a config file in a prod system.
const (
	grpcServerAddr    = "127.0.0.1:50000"
	grpcServerUseTLS  = false
	grpcServerCrtFile = "server.pem"
	grpcServerKeyFile = "server.key"
)

var (
	errs = make(chan error)
	sig  = make(chan os.Signal)
)

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
