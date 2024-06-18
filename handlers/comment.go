package handlers

import (
	"net/http"

	"blog-project/db"
	"blog-project/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// To create a new comment
func CreateComment(c *gin.Context) {
	var comment models.Comment
	// Attempts to bind the incoming JSON request body to the 'comment' model. 400
	if err := c.ShouldBindJSON(&comment); err != nil {
		logrus.Warn("Invalid comment data provided: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Attempts to create the comment in the database. 500
	if err := db.DB.Create(&comment).Error; err != nil {
		logrus.Error("Failed to create comment: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("Comment created successfully for blog entry ID: %d", comment.BlogID)
	c.JSON(http.StatusOK, comment)
}

func GetComment(c *gin.Context) {
	// Extract comment ID from URL parameter.
	commentID := c.Param("id")

	// Query database for comment with specified ID.
	var comment models.Comment
	// Attempts to fetch the comment from the database by ID. 404
	if err := db.DB.First(&comment, commentID).Error; err != nil {
		logrus.Warnf("Comment not found with ID: %s", commentID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Return the retrieved comment as JSON response.
	logrus.Infof("Comment retrieved successfully with ID: %s", commentID)
	c.JSON(http.StatusOK, comment)
}

// GetCommentsByBlogID retrieves comments for a specific blog entry ID.
func GetCommentsByBlogID(c *gin.Context) {
	blogID := c.Param("blog_id")
	var comments []models.Comment
	// Attempts to fetch comments from the database where blog_id matches the provided ID. 500
	if err := db.DB.Where("blog_id = ?", blogID).Find(&comments).Error; err != nil {
		logrus.Error("Failed to fetch comments for blog entry ID: ", blogID, "Error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}
	logrus.Infof("Comments retrieved successfully for blog entry ID: %s", blogID)
	c.JSON(http.StatusOK, comments)
}

// To retrieve all comments
func GetAllComments(c *gin.Context) {
	var comments []models.Comment
	// Attempts to fetch all comments from the database. 500
	if err := db.DB.Find(&comments).Error; err != nil {
		logrus.Errorf("Failed to fetch comments: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	logrus.Infof("Retrieved all comments: %v", comments)
	c.JSON(http.StatusOK, comments)
}
