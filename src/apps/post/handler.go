package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostHandler struct {
	db *gorm.DB
}

func NewPostHandler(db *gorm.DB) *PostHandler {
	return &PostHandler{db: db}
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	var posts []Post
	// Get all posts
	err := h.db.Find(&posts).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Get all posts failed",
			"data":    nil,
		})
		return
	}

	// Return all posts
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Get all posts success",
		"data":    posts,
	})
}

func (h *PostHandler) GetPost(c *gin.Context) {
	var post Post
	// Get post by id
	err := h.db.First(&post, c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Get post failed",
			"data":    nil,
		})
		return
	}

	// Return post
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Get post success",
		"data":    post,
	})
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var post Post
	// Bind JSON to post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Create post failed",
			"data":    nil,
		})
		return
	}

	// Create post
	err = h.db.Create(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Create post failed",
			"data":    nil,
		})
		return
	}

	// Return post
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Create post success",
		"data":    post,
	})
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	var post Post
	// Get post by id
	err := h.db.First(&post, c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Update post failed",
			"data":    nil,
		})
		return
	}

	// Bind JSON to post
	err = c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Update post failed",
			"data":    nil,
		})
		return
	}

	// Update post
	err = h.db.Save(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Update post failed",
			"data":    nil,
		})
		return
	}

	// Return post
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Update post success",
		"data":    post,
	})
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	var post Post
	// Get post by id
	err := h.db.First(&post, c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Delete post failed",
			"data":    nil,
		})
		return
	}

	// Delete post
	err = h.db.Delete(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Delete post failed",
			"data":    nil,
		})
		return
	}

	// Return post
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Delete post success",
		"data":    post,
	})
}
