package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	// workerID is the id assigned by the scheduler
	// after registering on scheduler.
	workerID string
)

func init() {
	loadConfig()
}

// Entry point of the worker application.
func main() {
	go registerWorker()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-sig:
		deregisterWorker()
		log.Printf("Signal (%d) received, stopping\n", s)
		os.Exit(0)
	}
}

func fatal(message string) {
	deregisterWorker()
	log.Fatalln(message)
}
