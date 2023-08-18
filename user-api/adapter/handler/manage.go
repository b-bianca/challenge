package manage

import (
	u "github.com/b-bianca/melichallenge/user-api/adapter/handler/controller"
	"github.com/b-bianca/melichallenge/user-api/internal/domain/port"
	"github.com/gin-gonic/gin"
)

type apps interface {
	RegisterRoutes(routes *gin.RouterGroup)
}

type Manage struct {
	user apps
}

type UseCases struct {
	User port.UserUseCase
}

func New(uc *UseCases) *Manage {

	userHandler := u.NewHandler(uc.User)

	return &Manage{
		user: userHandler,
	}
}

func (m *Manage) RegisterRoutes(group *gin.RouterGroup) {
	m.user.RegisterRoutes(group)
}
