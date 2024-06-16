package main

import (
	"blog-project/db"
	"blog-project/handlers"
	"blog-project/middleware"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Info("Logrus is running")

	f, _ := os.Create("logrus.log")

	multi := io.MultiWriter(f, os.Stdout)

	logrus.SetOutput(multi)

	db.Connect()
	r := gin.Default()

	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUser)
	r.GET("/users", handlers.GetAllUsers)

	r.POST("/login", handlers.Login)

	r.POST("/blog_entries", handlers.CreateBlogEntry)
	r.GET("/blog_entries/:id", handlers.GetBlogEntry)
	r.GET("/blog_entries", handlers.GetAllBlogEntries)

	r.GET("/comments", handlers.GetAllComments)

	authorized := r.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware())
	{
		authorized.POST("/blog_entries/:id/comments", handlers.CreateComment)
		authorized.GET("/blog_entries/:id/comments", handlers.GetCommentsByBlogID)
	}

	if err := r.Run(":8080"); err != nil {
		logrus.Fatalf("Failed to run the server: %v", err)
	}
}
