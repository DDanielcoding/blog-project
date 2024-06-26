package handlers

import (
	"blog-project/db"
	"blog-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateBlogEntry(c *gin.Context) {
	var blogEntry models.BlogEntry
	// Attempts to bind the incoming JSON request body to the blogEntry model.
	if err := c.ShouldBindJSON(&blogEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the user from the database based on some identifier (e.g., AuthorID from JWT token)
	var user models.User
	if err := db.DB.First(&user, blogEntry.AuthorID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find author"})
		return
	}

	// Set the Username field in the blog entry
	blogEntry.Username = user.Username

	// Set created_at and updated_at timestamps
	now := time.Now()
	blogEntry.CreatedAt = now
	blogEntry.UpdatedAt = now

	// Create the blog entry into the database.
	if err := db.DB.Create(&blogEntry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog entry"})
		return
	}

	c.JSON(http.StatusOK, blogEntry)
}

// To retrieve a blog entry by its ID.
func GetBlogEntry(c *gin.Context) {
	id := c.Param("id")
	var blogEntry models.BlogEntry
	//Attempts to fetch the blog entry from the database by ID.
	if err := db.DB.First(&blogEntry, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog entry not found"})
		return
	}
	c.JSON(http.StatusOK, blogEntry)
}

// To retrieve all blog entries.
func GetAllBlogEntries(c *gin.Context) {
	var blogEntries []models.BlogEntry
	// Attempts to fetch all blog entries from the database. '500'
	if err := db.DB.Find(&blogEntries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogEntries)
}
