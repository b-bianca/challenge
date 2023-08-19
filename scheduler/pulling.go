package scheduler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/b-bianca/melichallenge/notify-api/adapter/model"
)

func Pulling(url string, interval time.Duration) (notifications *model.NotificationListResponse, err error) {

	result, err := http.Get(url)
	if err != nil {
		return
	}

	defer result.Body.Close()

	err = json.NewDecoder(result.Body).Decode(&notifications)
	if err != nil {
		return nil, err
	}

	time.Sleep(interval)

	return notifications, nil
}
