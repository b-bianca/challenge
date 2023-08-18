package user

import (
	"github.com/b-bianca/melichallenge/internal/domain/port"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase port.UserUseCase
}

func NewHandler(u port.UserUseCase) *Handler {
	return &Handler{
		useCase: u,
	}
}

func (h *Handler) RegisterRoutes(routes *gin.RouterGroup) {
	userRoute := routes.Group("/user")
	userRoute.PATCH("/:id/", h.OptoutUser)
	userRoute.POST("/", h.CreateUser)
}
