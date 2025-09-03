package main

import (
	"todo-api/config"
	"todo-api/handlers"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

func main() {

	// Connect DB
	config.ConnectDatabase()

	// Run migrations
	config.DB.AutoMigrate(&models.Todo{})

	// Gin router
	r := gin.Default()

	// Test route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Todo routes
	r.POST("/todos", handlers.CreateTodo)
	r.GET("/todos", handlers.GetTodos)
	r.GET("/todos/:id", handlers.GetTodoByID)
	r.PUT("/todos/:id", handlers.UpdateTodo)
	r.DELETE("/todos/:id", handlers.DeleteTodo)
	// Start server
	r.Run(":8080")
}
