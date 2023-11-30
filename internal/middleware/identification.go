package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/blog/internal/response"
	"github.com/hixraid/blog/pkg/data/model"
	"github.com/hixraid/blog/pkg/service"
)

const (
	authHeader = "auth"
	userCtx    = "user_id"
)

func IdentifyUser(s service.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader(authHeader)
		if header == "" {
			response.NewErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerData := strings.Split(header, " ")
		if (len(headerData) != 2 || headerData[0] != "Bearer") || len(headerData[1]) == 0 {
			response.NewErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
			return
		}

		userId, err := s.ParseToken(headerData[1])
		if err != nil {
			response.NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		ctx.Set(userCtx, userId)
	}
}

func IdentifyAdmin(s service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := GetUserId(ctx)
		if err != nil {
			response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		user, err := s.GetById(userId)
		if err != nil {
			response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if user.Role != model.AdminRole {
			response.NewErrorResponse(ctx, http.StatusForbidden, err.Error())
			return
		}
	}
}

func GetUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(userCtx)
	if !ok {
		return -1, errors.New("not found user_id")
	}

	idInt, ok := id.(int)
	if !ok {
		return -1, errors.New("invalid type user_id")
	}

	return idInt, nil
}
