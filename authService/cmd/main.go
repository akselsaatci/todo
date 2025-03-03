package main

import (
	"authService/config"
	"authService/internal/adapters/handlers/Http"
	"authService/internal/adapters/repo"
	"authService/internal/core/services"
	"authService/internal/core/services/tokenService"
	"fmt"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
)

func main() {

	appConfig := config.ReadConfig()

	repoCredantials := repo.DbCredantials{
		Host:     appConfig.DATABASE_HOST,
		Port:     appConfig.DATABASE_PORT,
		User:     appConfig.DATABASE_USER,
		Password: appConfig.DATABASE_PASSWORD,
		DbName:   appConfig.DATABASE_NAME,
	}
	fmt.Println(os.Getenv("DB_HOST"))

	//jwtKey := []byte("26208abaa6dc35b2760ce14fa9c051ceca2caf08c67fcf79bbaa7f1177831b1e8c72114af2803fa9329d95cc6d859b187f87fdcea7d17f5a89389da82ae8e048")
	jwtKey := []byte(appConfig.JWT_SECRET_KEY)
	newTokenService := tokenService.NewTokenService(jwtKey)
	authRepo := repo.NewAuthPostgresRepository(repoCredantials)
	authService := services.NewAuthService(authRepo, newTokenService)

	router := gin.Default()
	authHandler := Http.NewHTTPHandler(*authService)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.POST("/auth", authHandler.AuthorizeUser)
	router.POST("/validate", authHandler.ValidateToken)
	router.POST("/register", authHandler.RegisterUser)
	err := router.Run(":" + appConfig.SERVER_PORT)
	if err != nil {
		return
	}
}
