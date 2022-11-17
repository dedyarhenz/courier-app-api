package main

import (
	"final-project-backend/config"
	"final-project-backend/handler"
	"final-project-backend/pkg/database/postgres"
	"final-project-backend/repository"
	"final-project-backend/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.InitConfig()
	db := postgres.New(cfg)

	userRepository := repository.NewUserRepositoryImpl(db)
	authUsecase := usecase.NewAuthUsecaseImpl(userRepository)
	authHandler := handler.NewAuthHandler(authUsecase)

	router := gin.Default()

	v1 := router.Group("v1")
	{
		auth := v1.Group("auth")
		{
			auth.GET("/login", authHandler.Login)
		}
	}

	router.Run(cfg.Server.Url)
}
