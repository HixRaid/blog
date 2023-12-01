package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/blog/internal/middleware"
	"github.com/hixraid/blog/internal/response"
	"github.com/hixraid/blog/pkg/data/model"
)

const userIdParam = "user_id"

// Summary: GetUserById;
// Tag: Users;
// Router: /api/users/:user_id [GET];
// Request: UserId;
// Response: UserOutput;
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

// Summary: GetAllUsers;
// Tag: Users;
// Router: /api/users [GET];
// Request: nil;
// Response: []UserOutput;
func (h *Handler) allUsers(ctx *gin.Context) {
	users, err := h.service.User.GetAll()
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	response.NewOkResponse(ctx, users)
}

// Summary: UpdateUser;
// Tag: Users;
// Router: /api/users [PUT];
// Request: Token;
// Response: Status;
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

// Summary: DeleteUser;
// Tag: Users;
// Router: /api/users [DELETE];
// Request: Token;
// Response: Status;
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
