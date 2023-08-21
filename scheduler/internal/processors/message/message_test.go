package message_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/b-bianca/melichallenge/scheduler/internal/processors/message"
	"github.com/b-bianca/melichallenge/scheduler/internal/processors/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	path = "/api/v1/notification/message"

	postMethod = "POST"
)

var (
	input = &model.MessageRequest{
		NotifyID: uuid.MustParse("d3caee6d-c806-4703-ae79-08bfc23fc79d"),
		Message:  "message",
	}

	messageJson = `{"notify_id":"d3caee6d-c806-4703-ae79-08bfc23fc79d","message":"message"}`
)

func TestSendMessage(t *testing.T) {
	t.Run("When api retuns status code saved; should return the expected message", func(t *testing.T) {

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, postMethod, r.Method)
			assert.Equal(t, path, r.URL.Path)

			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)
			assert.Equal(t, messageJson, string(body))

			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		webPageSender := &message.WebSender{
			WebURL: server.URL + path,
		}

		err := webPageSender.SendMessage(input)
		assert.NoError(t, err)

	})

	t.Run("When api retuns status code nok; should propagate the error to caller", func(t *testing.T) {

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, postMethod, r.Method)
			assert.Equal(t, path, r.URL.Path)

			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)
			assert.Equal(t, messageJson, string(body))

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}))
		defer server.Close()

		webSender := &message.WebSender{
			WebURL: server.URL + path,
		}

		err := webSender.SendMessage(input)
		assert.Error(t, err)
	})
}
