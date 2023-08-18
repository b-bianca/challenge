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
	user apps
}

type UseCases struct {
	Notify port.NotifyUseCase
}

func New(uc *UseCases) *Manage {

	notifyHandler := h.NewHandler(uc.Notify)

	return &Manage{
		user: notifyHandler,
	}
}

func (m *Manage) RegisterRoutes(group *gin.RouterGroup) {
	m.user.RegisterRoutes(group)
}
