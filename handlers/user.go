package handlers

import (
	"blog-project/db"
	"blog-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// To create a new user.
func CreateUser(c *gin.Context) {
	var user models.User
	// Attempts to bind the incoming JSON request body to the user mode. 400
	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.Errorf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Attempts to create the user in the database. 500
	if err := db.DB.Create(&user).Error; err != nil {
		logrus.Errorf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logrus.Infof("User created: %v", user)
	c.JSON(http.StatusOK, user)
}

//To retrieve a user by their ID.
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	// Attempts to fetch the user from the database by ID.
	if err := db.DB.First(&user, id).Error; err != nil {
		logrus.Errorf("User not found: %v")
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	logrus.Infof("User retrieved: %v", user)
	c.JSON(http.StatusOK, user)
}

// To retrieve all users.
func GetAllUsers(c *gin.Context) {
	var users []models.User
	// Attempts to fetch all users from the database. 500
	if err := db.DB.Find(&users).Error; err != nil {
		logrus.Errorf("Error fetching users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logrus.Infof("Users retrieved: %d users", len(users))
	c.JSON(http.StatusOK, users)
}
