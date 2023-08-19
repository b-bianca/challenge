package main

import (
	"log"

	"github.com/b-bianca/melichallenge/scheduler"
)

func main() {

	err := scheduler.Scheduler()
	if err != nil {
		log.Fatal("Erro no agendador:", err)
	}
}
