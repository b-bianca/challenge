package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/b-bianca/melichallenge/scheduler"
)

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go scheduler.Scheduler()

	<-interrupt

}
