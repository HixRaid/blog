package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
			auth.PUT("/", h.updateUser)
			auth.DELETE("/", h.deleteUser)
		}

		posts := api.Group("/posts")
		{
			posts.POST("/", h.createPost)
			posts.GET("/", h.getAllPosts)
			posts.GET("/:id", h.getPostById)
			posts.PUT("/:id", h.updatePostById)
			posts.DELETE("/:id", h.deletePostById)

			comments := api.Group("/:id/comments")
			{
				comments.POST("/", h.createComment)
				comments.GET("/", h.getCommentsByPostId)
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
