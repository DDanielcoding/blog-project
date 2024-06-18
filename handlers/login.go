package handlers

import (
	"blog-project/db"
	"blog-project/middleware"
	"blog-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)
// To authenticate a user and provide a JWT token if authentication is successful.
func Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	logrus.Info("Login attempt started")
	// Attempts to bind the incoming JSON request body to the loginData struct. 400
	if err := c.ShouldBindJSON(&loginData); err != nil {
		logrus.Warn("Invalid login data provided: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Info("Login data received: %v", loginData)

	var user models.User
	// Attempts to fetch the user from the database by email. 401
	if err := db.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		logrus.Warn("Invalid email: ", loginData.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}
	// Compares the provided password with the stored password. 401
	if loginData.Password != user.Password {
		logrus.Warn("Invalid password for email: ", loginData.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}
	// Call middleware.GenerateToken to create a JWT token for the authenticated user. 500
	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		logrus.Error("Failed to generate token: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	logrus.Infof("Token generated successfully for user ID: %d", user.ID)

	c.JSON(http.StatusOK, gin.H{"token": token})
	logrus.Info("Login attempt successful")
}
