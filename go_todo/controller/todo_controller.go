package controller

import (
	"go-todo/models"
	"go-todo/service"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	todos, err := service.GetAllTodos()
	if err != nil {
		c.JSON(500, gin.H{
			"succes": false,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "Here are the todos",
		"data":    todos,
	})
}

func CreateTodo(c *gin.Context) {
	var input models.Todo
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	todo, err := service.CreateTodo(input)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"success": true,
		"message": "Todo created successfully",
		"data":    todo,
	})
	
}