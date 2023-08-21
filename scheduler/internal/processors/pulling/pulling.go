package pulling

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/b-bianca/melichallenge/scheduler/internal/processors/model"
)

func Pulling(url string) (notification *model.NotificationList, err error) {

	time.Sleep(20 * time.Second)
	result, err := http.Get(url)
	if err != nil {
		return
	}

	defer result.Body.Close()

	err = json.NewDecoder(result.Body).Decode(&notification)
	if err != nil {
		return nil, err
	}

	return notification, nil
}
