package db

import (
	"blog-project/models"

	"github.com/sirupsen/logrus"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("blog_project.db"), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("Failed to connect to the database: %v", err)
	}

	logrus.Info("Database connected successfully!")

	err = DB.AutoMigrate(&models.User{}, &models.BlogEntry{}, &models.Comment{})
	if err != nil {
		logrus.Fatalf("Failed to migrate database: %v", err)
	}

	logrus.Info("Database migration completed successfully")
}
