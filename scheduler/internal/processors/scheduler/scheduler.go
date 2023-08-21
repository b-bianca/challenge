package scheduler

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/b-bianca/melichallenge/scheduler/internal/processors/message"
	"github.com/b-bianca/melichallenge/scheduler/internal/processors/model"
	"github.com/b-bianca/melichallenge/scheduler/internal/processors/pulling"
	_ "github.com/joho/godotenv/autoload"
)

var (
	notifyAPIURL = os.Getenv("NOTIFY_API_URL")
)

type Scheduler struct {
	WebSender message.MessageSender
}

func New(w message.MessageSender) *Scheduler {
	return &Scheduler{
		WebSender: w,
	}
}

func (s *Scheduler) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			notification, err := pulling.Pulling(notifyAPIURL)
			fmt.Println(notification, "not")
			if err != nil {
				log.Println("Error fetching notifications", err)
			} else {
				now := time.Now()

				for _, n := range notification.Result {
					timeDifference := n.DateTime.Sub(now)
					if timeDifference > 0 {
						go func(n *model.Notification) {
							defer func() {

							}()

							timer := time.NewTimer(timeDifference)

							select {
							case <-ctx.Done():
								timer.Stop()
								return
							case <-timer.C:
								body := &model.MessageRequest{
									NotifyID: n.ID,
									Message:  n.Message,
								}

								err := message.SendMessage(s.WebSender, body)
								if err != nil {
									log.Println("Error sending message", err)
								} else {
									fmt.Println("Message sent!", body.Message)
								}
							}
						}(n)
					}
				}
			}
		}
	}
}
