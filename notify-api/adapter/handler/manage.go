package manage

import (
	h "github.com/b-bianca/melichallenge/notify-api/adapter/handler/controller"
	"github.com/b-bianca/melichallenge/notify-api/internal/domain/port"
	"github.com/gin-gonic/gin"
)

type apps interface {
	RegisterRoutes(routes *gin.RouterGroup)
}

type Manage struct {
	notify apps
}

type UseCases struct {
	Notify port.NotifyUseCase
}

func New(uc *UseCases) *Manage {

	notifyHandler := h.NewHandler(uc.Notify)

	return &Manage{
		notify: notifyHandler,
	}
}

func (m *Manage) RegisterRoutes(group *gin.RouterGroup) {
	m.notify.RegisterRoutes(group)
}
