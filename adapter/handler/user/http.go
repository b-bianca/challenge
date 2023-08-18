package user

import (
	"net/http"

	"github.com/b-bianca/melichallenge/adapter/model"
	"github.com/b-bianca/melichallenge/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserParam struct {
	ID string `uri:"id" binding:"required"`
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	var input model.UserRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	domain := &entity.User{
		Name:     input.Name,
		CPF:      input.CPF,
		Email:    input.Email,
		Password: input.Password,
	}

	res, err := h.useCase.CreateUser(ctx, domain)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	output := &model.UserResponse{
		ID:           res.ID,
		Name:         res.Name,
		Email:        res.Email,
		CPF:          res.CPF,
		Notification: true,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, output)
}

func (h *Handler) OptoutUser(ctx *gin.Context) {
	idParam := ctx.Param("id")

	var input model.OptoutRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	domain := &entity.User{
		ID:           uuid.MustParse(idParam),
		Notification: false,
	}

	err := h.useCase.OptoutUser(ctx, domain)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	ctx.JSON(http.StatusNoContent, "")
}
