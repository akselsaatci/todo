package main

import (
	"github.com/gin-gonic/gin"
	"toDoService/internal/adapters/handlers/ApiHandlers"
	"toDoService/internal/adapters/handlers/CommunicationHandlers/Http"
	"toDoService/internal/adapters/repo"
	"toDoService/internal/core/services"
)

func main() {

	toDoRepo := repo.NewTodoPostgresRepository()
	comHandler := Http.NewHttpCommunicationHandler()
	todoService := services.NewTodoService(toDoRepo, comHandler)

	router := gin.Default()
	todoHandler := ApiHandlers.HTTPHandler{todoService}
	router.POST("/todo", todoHandler.GetUserTodos) // im gonna change this to get
	err := router.Run(":3000")
	if err != nil {
		return
	}
}
