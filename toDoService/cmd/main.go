package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"toDoService/internal/adapters/handlers/ApiHandlers"
	"toDoService/internal/adapters/handlers/CommunicationHandlers/Http"
	"toDoService/internal/adapters/repo"
	"toDoService/internal/core/services"
)

func main() {

	repoCredantials := repo.DbCredantials{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	}
	authUrl := os.Getenv("AUTH_SERVICE")

	toDoRepo := repo.NewTodoPostgresRepository(repoCredantials)
	comHandler := Http.NewHttpCommunicationHandler(authUrl)
	todoService := services.NewTodoService(toDoRepo, comHandler)

	router := gin.Default()
	todoHandler := ApiHandlers.HTTPHandler{todoService}
	router.GET("/todo", todoHandler.GetUserTodos)
	router.POST("/todo", todoHandler.AddTodo)
	router.PATCH("/todo", todoHandler.UpdateTodo)
	router.DELETE("/todo", todoHandler.DeleteTodo)
	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		return
	}
}
