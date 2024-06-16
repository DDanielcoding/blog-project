package main

import (
	"blog-project/db"
	"blog-project/handlers"
	"blog-project/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.Default()

	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUser)

	r.POST("/login", handlers.Login)

	r.POST("/blog_entries", handlers.CreateBlogEntry)
	r.GET("/blog_entries/:id", handlers.GetBlogEntry)
	r.GET("/blog_entries", handlers.GetAllBlogEntries)

	authorized := r.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware())
	{
		authorized.POST("/blog_entries/:id/comments", handlers.CreateComment)
		authorized.GET("/blog_entries/:id/comments", handlers.GetCommentsByBlogID)
	}

	r.Run(":8080")
}
