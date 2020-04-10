package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmlwo530/maru-web-app-project/db/models"
	"github.com/jinzhu/gorm"
)

type createPostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type updatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// FindPosts is get all posts
func FindPosts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var posts []models.Post
	db.Find(&posts)

	c.JSON(http.StatusOK, posts)
}

// CreatePost is create a post
func CreatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input createPostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Title: input.Title, Content: input.Content}

	db.Create(&post)

	c.JSON(http.StatusOK, post)
}

// FindPost is find a post
func FindPost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var post models.Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// UpdatePost is update a post
func UpdatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var post models.Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input updatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&post).Updates(input)

	c.JSON(http.StatusOK, post)
}

// DeletePost is delete a post
func DeletePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var post models.Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
