package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/blog/internal/response"
	"github.com/hixraid/blog/pkg/data/model"
)

type signInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type tokenResponse struct {
	Token string `json:"token"`
}

// Summary: SignUp;
// Tag: Auth;
// Router: /auth/sign-up [POST];
// Request: UserInput;
// Response: UserId;
func (h *Handler) signUp(ctx *gin.Context) {
	var input model.UserInput
	if err := ctx.Bind(&input); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	userId, err := h.service.Auth.CreateUser(input)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewOkResponse(ctx, model.UserIdResponse{
		UserId: userId,
	})
}

// Summary: SignIn;
// Tag: Auth;
// Router: /auth/sign-in [POST];
// Request: SignInInput;
// Response: Token;
func (h *Handler) signIn(ctx *gin.Context) {
	var input signInInput
	if err := ctx.Bind(&input); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.service.Auth.GenerateToken(input.Email, input.Password)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewOkResponse(ctx, tokenResponse{
		Token: token,
	})
}
