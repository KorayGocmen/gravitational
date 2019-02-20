package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	loadConfig()
}

// Entry point of the scheduler application.
func main() {

	go startGRPCServer()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-sig:
		log.Printf("Signal (%d) received, stopping\n", s)
		os.Exit(0)
	}
}
