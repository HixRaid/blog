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
		users := api.Group("/users")
		{
			users.POST("/sign-up", h.signUp)
			users.POST("/sign-in", h.signIn)
			users.GET("/:id", h.getUserById)
			users.GET("/", h.getAllUsers)
			users.PUT("/", h.updateUser)
			users.DELETE("/", h.deleteUser)

			posts := api.Group("/:id/posts")
			{
				posts.POST("/", h.createPost)
				posts.GET("/", h.getPostsByUserId)
			}
		}

		posts := api.Group("/posts")
		{
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
			comments.GET("/:id", h.getCommentById)
			comments.PUT("/:id", h.updateCommentById)
			comments.DELETE("/:id", h.deleteCommentById)
		}
	}

	return router
}
