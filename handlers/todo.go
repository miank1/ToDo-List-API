package handlers

import (
	"fmt"
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

// GetTodos handles GET /todos
func GetTodos(c *gin.Context) {

	page := 1
	limit := 10

	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page) // ??
	}

	if l := c.Query("limit"); l != "" {
		fmt.Sscanf(l, "%d", &limit) // ??
	}

	// Calculate offset
	offset := (page - 1) * limit

	var todos []models.Todo
	result := config.DB.Limit(limit).Offset(offset).Find(&todos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":  page,
		"limit": limit,
		"data":  todos,
	})
}

// GetTodoByID handles GET /todos/:id
func GetTodoByID(c *gin.Context) {
	id := c.Param("id") // get id from URL

	var todo models.Todo

	// Look for todo with given id
	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// UpdateTodo handles PUT /todos/:id
func UpdateTodo(c *gin.Context) {
	id := c.Param("id") // get todo id from URL
	var todo models.Todo

	// Find todo
	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// Bind JSON body into existing todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save changes
	if err := config.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo handles DELETE /todos/:id
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	// Check if todo exists
	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// Delete todo
	if err := config.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
