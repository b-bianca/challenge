package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/b-bianca/melichallenge/scheduler/internal/processors/message"
	"github.com/b-bianca/melichallenge/scheduler/internal/processors/scheduler"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	ctx, _ := context.WithCancel(context.Background())

	messageWeb := message.NewWebSender(os.Getenv("WEB_API"))

	scheduler := scheduler.New(messageWeb)

	scheduler.Start(ctx)

	<-interrupt

}
