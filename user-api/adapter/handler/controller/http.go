package controller

import (
	"log"
	"net/http"

	"github.com/b-bianca/melichallenge/user-api/adapter/model"
	"github.com/b-bianca/melichallenge/user-api/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateUser(ctx *gin.Context) {
	var input model.UserRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Println(
			`"event": "deserialize_error"`,
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	domain := &entity.User{
		CPF:          input.CPF,
		Notification: true,
	}

	res, err := h.useCase.CreateUser(ctx, domain)
	if err != nil {
		log.Println(
			`"event": "use_case_create_failed"`,
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	output := &model.UserResponse{
		ID:           res.ID,
		CPF:          res.CPF,
		Notification: true,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, output)
}

func (h *Handler) PartialUpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")

	var input model.OptoutRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Println(
			`"event": "deserialize_error"`,
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	domain := &entity.User{
		ID:           uuid.MustParse(idParam),
		Notification: input.Notification,
	}

	err := h.useCase.PartialUpdateUser(ctx, domain)
	if err != nil {
		log.Println(
			`"event": "user_case_updated_user_failed"`,
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	ctx.JSON(http.StatusNoContent, "")
}
