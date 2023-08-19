package scheduler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/b-bianca/melichallenge/notify-api/adapter/model"
)

type MessageSender interface {
	SendMessage(*model.MessageRequest) error
}

type WebPageSender struct {
	WebPageURL string
}

func (w *WebPageSender) SendMessage(body *model.MessageRequest) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", w.WebPageURL, bytes.NewReader(b))
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

func SendMessage(sender MessageSender, message *model.MessageRequest) {
	err := sender.SendMessage(message)
	if err != nil {
		fmt.Println("Erro ao enviar mensagem:", err)
	}
}
