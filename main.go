package main

import (
	"todo-api/config"
	"todo-api/handlers"
	"todo-api/middleware"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

func main() {

	// Connect DB
	config.ConnectDatabase()

	// Run migrations
	config.DB.AutoMigrate(&models.Todo{}, &models.User{})

	// Gin router
	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Test route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	todoRoutes := r.Group("/todos")
	todoRoutes.Use(middleware.AuthMiddleware())
	{
		todoRoutes.POST("", handlers.CreateTodo)
		todoRoutes.GET("", handlers.GetTodos)
		todoRoutes.GET("/:id", handlers.GetTodoByID)
		todoRoutes.PUT("/:id", handlers.UpdateTodo)
		todoRoutes.DELETE("/:id", handlers.DeleteTodo)
	}
	// Start server
	r.Run(":8080")
}
