package scheduler

import (
	"sync"
	"time"

	"github.com/b-bianca/melichallenge/notify-api/adapter/model"
)

const (
	notifyAPIURL = "http://localhost:8081/api/v1/notification"
)

func Scheduler() error {

	webPageSender := &WebPageSender{WebPageURL: "http://localhost:8081/api/v1/notification/message"}

	notification, err := Pulling(notifyAPIURL, 1)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, n := range notification.Result {

		body := &model.MessageRequest{
			NotifyID: n.ID,
			Message:  n.Message,
		}

		now := time.Now()
		duration := n.DateTime.Sub(now)
		if duration > 0 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				time.Sleep(duration)
				SendMessage(webPageSender, body)
			}()

		}

	}
	wg.Wait()
	return nil
}
