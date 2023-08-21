package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/b-bianca/melichallenge/scheduler/internal/processors/model"
)

type MessageSender interface {
	SendMessage(*model.MessageRequest) error
}

type WebSender struct {
	WebURL string
}

func NewWebSender(WebURL string) *WebSender {
	return &WebSender{
		WebURL: WebURL,
	}
}

func (w *WebSender) SendMessage(body *model.MessageRequest) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", w.WebURL, bytes.NewReader(b))
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func SendMessage(sender MessageSender, message *model.MessageRequest) error {
	err := sender.SendMessage(message)
	if err != nil {
		fmt.Println("Failed to send message", err)
		return err
	}

	return nil
}
