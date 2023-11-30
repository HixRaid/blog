package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hixraid/blog/internal/middleware"
	"github.com/hixraid/blog/pkg/service"
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
			users.GET("/:user_id", h.userById)
			users.GET("/", h.allUsers)
			users.PUT("/", middleware.IdentifyUser(h.service.Auth), h.updateUser)
			users.DELETE("/", middleware.IdentifyUser(h.service.Auth), h.deleteUser)
		}

		posts := api.Group("/posts")
		{
			posts.POST("/", h.createPost, middleware.IdentifyUser(h.service.Auth), middleware.IdentifyAdmin(h.service.User))
			posts.GET("/", h.allPosts)
			posts.GET("/:post_id", h.postById)
			posts.PUT("/:post_id", middleware.IdentifyUser(h.service.Auth), middleware.IdentifyAdmin(h.service.User), h.updatePostById)
			posts.DELETE("/:post_id", middleware.IdentifyUser(h.service.Auth), middleware.IdentifyAdmin(h.service.User), h.deletePostById)

			comments := api.Group("/:post_id/comments")
			{
				comments.POST("/", middleware.IdentifyUser(h.service.Auth), middleware.IdentifyUser(h.service.Auth), h.createComment)
				comments.GET("/", h.commentsByPostId)
			}
		}

		comments := api.Group("/comments")
		{
			comments.PUT("/:comment_id", middleware.IdentifyUser(h.service.Auth), h.updateCommentById)
			comments.DELETE("/:comment_id", middleware.IdentifyUser(h.service.Auth), h.deleteCommentById)
		}
	}

	return router
}
