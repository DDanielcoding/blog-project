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

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{

	}

	r.Run(":8080")
}
