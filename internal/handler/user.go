package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/blog/internal/data/model"
	"github.com/hixraid/blog/internal/middleware"
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
	userId, err := middleware.GetUserId(ctx)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var input model.UserInput
	if err := ctx.Bind(&input); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	if err := h.service.User.UpdateById(userId, input); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.NewStatusResponse(ctx, "OK")
}

func (h *Handler) deleteUser(ctx *gin.Context) {
	userId, err := middleware.GetUserId(ctx)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.service.User.DeleteById(userId); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.NewStatusResponse(ctx, "OK")
}
