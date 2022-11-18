package main

import (
	"final-project-backend/config"
	"final-project-backend/handler"
	"final-project-backend/middleware"
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

	addressRepository := repository.NewAddressRepositoryImpl(db)
	addressUsecase := usecase.NewAddressUsecaseImpl(addressRepository)
	addressHandler := handler.NewAddressHandler(addressUsecase)

	router := gin.Default()

	v1 := router.Group("v1")
	{
		auth := v1.Group("auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
		}

		admin := v1.Group("admin")
		{
			addresses := admin.Group("addresses")
			{
				addresses.GET("/", addressHandler.GetAllAddress)
			}
		}

		users := v1.Group("users")
		{
			users.Use(middleware.UserAccess())

			users.GET("/addresses", addressHandler.GetAddressByUser)
			users.POST("/addresses", addressHandler.CreateAddress)
		}
	}

	router.Run(cfg.Server.Url)
}
