package pulling_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/b-bianca/melichallenge/scheduler/internal/processors/model"
	"github.com/b-bianca/melichallenge/scheduler/internal/processors/pulling"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	notifications = &model.NotificationList{
		Result: []*model.Notification{
			{
				ID:      uuid.MustParse("d3caee6d-c806-4703-ae79-08bfc23fc79d"),
				Message: "message-test-1"},
			{
				ID:      uuid.MustParse("d3caee6d-c806-4703-ae79-08bfc23fc79d"),
				Message: "message-test-2"},
		},
	}
)

func TestPulling(t *testing.T) {
	t.Run("When everytghing goes ok, should return notification list", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(notifications)
		}))
		defer server.Close()

		notifications, err := pulling.Pulling(server.URL)
		assert.NoError(t, err)
		assert.NotNil(t, notifications)
	})
	t.Run("When external api return error, should return an error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)

		}))
		defer server.Close()

		notifications, err := pulling.Pulling(server.URL)
		assert.Error(t, err)
		assert.Nil(t, notifications)
	})

}
