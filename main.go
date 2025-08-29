package main

import (
	"github.com/example/internal/handlers"
	"gofr.dev/pkg/gofr"
)

func main() {

	app := gofr.New()

	app.GET("/ping", func(c *gofr.Context) (any, error) {
		return "pong", nil
	})

	app.GET("/posts", handlers.NewPostHandler().GetPosts)
	app.GET("/posts/{id}", handlers.NewPostHandler().GetPostById)
	app.POST("/posts", handlers.NewPostHandler().CreatePost)
	app.PUT("/posts/{id}", handlers.NewPostHandler().EditPost)
	app.PATCH("/posts/{id}", handlers.NewPostHandler().EditPostByFields)
	app.DELETE("/posts/{id}", handlers.NewPostHandler().DeletePost)

	app.POST("/comments", handlers.NewCommentHandler().CreateComment)
	app.GET("/posts/{id}/comments", handlers.NewCommentHandler().GetCommentsByPostId)
	app.GET("/comments", handlers.NewCommentHandler().GetComments)

	app.Run()

}
