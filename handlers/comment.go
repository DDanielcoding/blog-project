// handlers/comment.go

package handlers

import (
	"net/http"

	"blog-project/db"
	"blog-project/models"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func GetComment(c *gin.Context) {
	// Extract comment ID from URL parameter
	commentID := c.Param("id")

	// Query database for comment with specified ID
	var comment models.Comment
	if err := db.DB.First(&comment, commentID).Error; err != nil {
		// Handle case where comment is not found
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Return the retrieved comment as JSON response
	c.JSON(http.StatusOK, comment)
}

// GetCommentsByBlogID retrieves comments for a specific blog entry ID
func GetCommentsByBlogID(c *gin.Context) {
	blogID := c.Param("blog_id")
	var comments []models.Comment
	if err := db.DB.Where("blog_id = ?", blogID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}
	c.JSON(http.StatusOK, comments)
}
