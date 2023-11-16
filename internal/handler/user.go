package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const userIdParam = "user_id"

func (h *Handler) userById(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param(userIdParam))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid param user_id")
		return
	}

	user, err := h.service.User.GetById(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) allUsers(ctx *gin.Context) {
	users, err := h.service.User.GetAll()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *Handler) updateUser(ctx *gin.Context) {
}

func (h *Handler) deleteUser(ctx *gin.Context) {
}
