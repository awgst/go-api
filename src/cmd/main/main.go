package main

import (
	"go-api/src/apps/post"
	"go-api/src/config/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	db := database.Connect()

	// Migrate database
	db.AutoMigrate(&post.Post{})

	// Create post handler
	postHandler := post.NewPostHandler(db)

	// Create router
	router := gin.Default()
	api := router.Group("/api")
	postRoute := api.Group("/post")
	postRoute.GET("", postHandler.GetPosts)
	postRoute.GET("/:id", postHandler.GetPost)
	postRoute.POST("", postHandler.CreatePost)
	postRoute.PUT("/:id", postHandler.UpdatePost)
	postRoute.DELETE("/:id", postHandler.DeletePost)

	// Run server
	router.Run()
}
