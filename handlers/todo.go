package handlers

import (
	"net/http"
	"todo-api/config"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

// CreateTodo handles POST /todos
func CreateTodo(c *gin.Context) {
	var todo models.Todo

	// Bind JSON input to struct
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Default status = "pending" if not provided
	if todo.Status == "" {
		todo.Status = "pending"
	}

	// Save to DB
	if err := config.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}
