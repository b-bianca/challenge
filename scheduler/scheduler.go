package scheduler

import (
	"fmt"
	"time"

	"github.com/b-bianca/melichallenge/notify-api/adapter/model"
	"github.com/go-co-op/gocron"
)

const (
	notifyAPIURL = "http://localhost:8081/api/v1/notification"
)

func Scheduler() error {

	taskCompleted := make(chan struct{})

	s := gocron.NewScheduler(time.UTC)

	notification, err := Pulling(notifyAPIURL, 1)
	if err != nil {
		return err
	}

	webPageSender := &WebPageSender{WebPageURL: "http://localhost:8081/api/v1/notification/message"}

	for _, n := range notification.Result {

		isoDateTime := n.DateTime.Format("2006-01-02T15:04:05.999999999-07:00")

		dbTime, err := time.Parse("2006-01-02T15:04:05.999999999-07:00", isoDateTime)
		if err != nil {
			fmt.Println("Erro ao fazer parsing do horário:", err)
			return err
		}

		body := &model.MessageRequest{
			NotifyID: n.ID,
			Message:  n.Message,
		}
		s.At(dbTime).Do(func() {
			fmt.Println("Tarefa agendada para:", dbTime)
			SendMessage(webPageSender, body)
			taskCompleted <- struct{}{}
		})
	}

	s.StartAsync()

	go func() {
		ticker := time.NewTicker(time.Second * 5)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				for i := 0; i < len(notification.Result); i++ {
					<-taskCompleted
				}
				fmt.Println("Tarefas agendadas concluídas!")
			}
		}
	}()

	select {}
}
