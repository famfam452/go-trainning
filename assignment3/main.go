package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"tanakrit.karaket/assignment3/handlers"
	custom_middleware "tanakrit.karaket/assignment3/middlewere"
)

func main() {
	echoServer := echo.New()

	echoServer.Use(custom_middleware.ZerologMiddleware)

	echoServer.POST("/tasks", handlers.CreateTask)
	echoServer.GET("/tasks", handlers.GetAllTasks)
	echoServer.GET("/tasks/:id", handlers.GetTaskById)
	echoServer.PATCH("/tasks/:id", handlers.PatchTaskById)
	echoServer.DELETE("tasks/:id", handlers.DeleteTask)

	if err := echoServer.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
