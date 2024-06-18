package main

import (
	"blog-project/db"
	"blog-project/handlers"
	"blog-project/middleware"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Info("Logrus is running")

	// Generate a log file name with a timestamp
	logFileName := fmt.Sprintf("logs/logrus_%s.log", time.Now().Format("20060102_150405"))

	f, err := os.Create(logFileName)
	if err != nil {
		logrus.Fatalf("Failed to create log file: %v", err)
	}
	// Writes logs to both the file and standard output.
	multi := io.MultiWriter(f, os.Stdout)

	logrus.SetOutput(multi)
	// Calls the Connect function from the db package to initialize and connect to the database.
	db.Connect()
	// Initializes a Gin router with default middleware.
	r := gin.Default()

	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUser)
	r.GET("/users", handlers.GetAllUsers)

	r.POST("/login", handlers.Login)

	r.POST("/blog_entries", handlers.CreateBlogEntry)
	r.GET("/blog_entries/:id", handlers.GetBlogEntry)
	r.GET("/blog_entries", handlers.GetAllBlogEntries)

	r.GET("/comments", handlers.GetAllComments)
	r.POST("/blog_entries/:id/comments", handlers.CreateComment)
	r.GET("/blog_entries/:id/comments", handlers.GetCommentsByBlogID)

	authorized := r.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware())
	{

	}
	// Starts the Gin server on port 8080.
	if err := r.Run(":8080"); err != nil {
		logrus.Fatalf("Failed to run the server: %v", err)
	}
}
