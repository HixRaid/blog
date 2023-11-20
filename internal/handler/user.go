package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/blog/internal/response"
)

const userIdParam = "user_id"

func (h *Handler) userById(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param(userIdParam))
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid param user_id")
		return
	}

	user, err := h.service.User.GetById(userId)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewOkResponse(ctx, user)
}

func (h *Handler) allUsers(ctx *gin.Context) {
	users, err := h.service.User.GetAll()
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	response.NewOkResponse(ctx, users)
}

func (h *Handler) updateUser(ctx *gin.Context) {
}

func (h *Handler) deleteUser(ctx *gin.Context) {
}
