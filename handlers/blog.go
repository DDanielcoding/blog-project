package handlers

import (
	"blog-project/db"
	"blog-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBlogEntry(c *gin.Context) {
	var blogEntry models.BlogEntry
	if err := c.ShouldBindJSON(&blogEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&blogEntry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogEntry)
}

func GetBlogEntry(c *gin.Context) {
	id := c.Param("id")
	var blogEntry models.BlogEntry
	if err := db.DB.First(&blogEntry, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog entry not found"})
		return
	}
	c.JSON(http.StatusOK, blogEntry)
}

func GetAllBlogEntries(c *gin.Context) {
	var blogEntries []models.BlogEntry
	if err := db.DB.Find(&blogEntries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogEntries)
}
