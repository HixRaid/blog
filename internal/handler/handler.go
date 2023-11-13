package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hixraid/blog/internal/service"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/:id", h.userById)
			users.GET("/", h.allUsers)
			users.PUT("/", h.updateUser)
			users.DELETE("/", h.deleteUser)
		}

		posts := api.Group("/posts")
		{
			posts.POST("/", h.createPost)
			posts.GET("/", h.allPosts)
			posts.GET("/:id", h.postById)
			posts.PUT("/:id", h.updatePostById)
			posts.DELETE("/:id", h.deletePostById)

			comments := api.Group("/:id/comments")
			{
				comments.POST("/", h.createComment)
				comments.GET("/", h.commentsByPostId)
			}
		}

		comments := api.Group("/comments")
		{
			comments.PUT("/:id", h.updateCommentById)
			comments.DELETE("/:id", h.deleteCommentById)
		}
	}

	return router
}
