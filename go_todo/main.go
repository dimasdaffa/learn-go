package main

import (
	// "fmt"
	"go-todo/config"
	"go-todo/controller"

	"github.com/gin-gonic/gin"
)

func main ()  {
	config.ConnectDatabase()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/todos", controller.GetTodos) //API Endpoint
	r.POST("/todos", controller.CreateTodo) //API Endpoint

	r.Run() // listen and serve on 0.0.0.0:8080
}