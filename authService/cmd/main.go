package main

import (
	"authService/internal/adapters/handlers/Http"
	"authService/internal/adapters/repo"
	"authService/internal/core/services"
	"authService/internal/core/services/tokenService"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	repoCredantials := repo.DbCredantials{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	}

	//jwtKey := []byte("26208abaa6dc35b2760ce14fa9c051ceca2caf08c67fcf79bbaa7f1177831b1e8c72114af2803fa9329d95cc6d859b187f87fdcea7d17f5a89389da82ae8e048")
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	newTokenService := tokenService.NewTokenService(jwtKey)
	authRepo := repo.NewAuthPostgresRepository(repoCredantials)
	authService := services.NewAuthService(authRepo, newTokenService)

	router := gin.Default()
	authHandler := Http.NewHTTPHandler(*authService)
	router.POST("/auth", authHandler.AuthorizeUser)
	router.POST("/validate", authHandler.ValidateToken)
	router.POST("/register", authHandler.RegisterUser)
	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		return
	}
}
