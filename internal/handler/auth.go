package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/blog/internal/data/model"
)

type signInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type tokenResponse struct {
	Token string `json:"token"`
}

func (h *Handler) signUp(ctx *gin.Context) {
	var input model.UserInput
	if err := ctx.Bind(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	userId, err := h.service.Auth.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, model.UserIdResponse{
		UserId: userId,
	})
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input signInInput
	if err := ctx.Bind(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.service.Auth.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, tokenResponse{
		Token: token,
	})
}
