package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type statusResponse struct {
	Status string `json:"status"`
}

func NewStatusResponse(ctx *gin.Context, message string) {
	NewOkResponse(ctx, statusResponse{message})
}

func NewOkResponse(ctx *gin.Context, obj any) {
	ctx.JSON(http.StatusOK, obj)
}
