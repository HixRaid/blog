package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/blog/internal/data/model"
	"github.com/hixraid/blog/internal/response"
)

const postIdParam = "post_id"

func (h *Handler) createPost(ctx *gin.Context) {
	var input model.PostInput
	if err := ctx.Bind(&input); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	postId, err := h.service.Post.Create(input)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewOkResponse(ctx, model.PostIdResponse{
		PostId: postId,
	})
}

func (h *Handler) allPosts(ctx *gin.Context) {
	posts, err := h.service.Post.GetAll()
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewOkResponse(ctx, posts)
}

func (h *Handler) postById(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param(postIdParam))
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid param post_id")
		return
	}

	post, err := h.service.Post.GetById(postId)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewOkResponse(ctx, post)
}

func (h *Handler) updatePostById(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param(postIdParam))
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid param post_id")
		return
	}

	var input model.PostInput
	if err := ctx.Bind(&input); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	if err = h.service.Post.UpdateById(postId, input); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewStatusResponse(ctx, "OK")
}

func (h *Handler) deletePostById(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param(postIdParam))
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid param post_id")
		return
	}

	if err = h.service.Post.DeleteById(postId); err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewStatusResponse(ctx, "OK")
}
