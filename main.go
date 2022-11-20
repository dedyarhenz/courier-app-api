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
	userUsecase := usecase.NewUserUsecaseImpl(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	authUsecase := usecase.NewAuthUsecaseImpl(userRepository)
	authHandler := handler.NewAuthHandler(authUsecase)

	addressRepository := repository.NewAddressRepositoryImpl(db)
	addressUsecase := usecase.NewAddressUsecaseImpl(addressRepository)
	addressHandler := handler.NewAddressHandler(addressUsecase)

	sizeRepository := repository.NewSizeRepositoryImpl(db)
	sizeUsecase := usecase.NewSizeUsecaseImpl(sizeRepository)
	sizeHandler := handler.NewSizeHandler(sizeUsecase)

	categoryRepository := repository.NewCategoryRepositoryImpl(db)
	categoryUsecase := usecase.NewCategoryUsecaseImpl(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	addOnRepository := repository.NewAddOnRepositoryImpl(db)
	addOnUsecase := usecase.NewAddOnUsecaseImpl(addOnRepository)
	addOnHandler := handler.NewAddOnHandler(addOnUsecase)

	router := gin.Default()

	v1 := router.Group("v1")
	{
		auth := v1.Group("auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
		}

		addresses := v1.Group("addresses")
		{
			addresses.Use(middleware.CheckAuth(), middleware.AdminAccess())
			addresses.GET("/", addressHandler.GetAllAddress)
		}

		users := v1.Group("users")
		{
			users.Use(middleware.CheckAuth(), middleware.UserAccess())
			users.GET("/", userHandler.GetUserById)
			users.POST("/top-up", userHandler.TopUp)

			addresses := users.Group("addresses")
			{
				addresses.GET("/", addressHandler.GetAddressByUser)
				addresses.POST("/", addressHandler.CreateAddress)
			}
		}

		sizes := v1.Group("sizes")
		{
			sizes.Use(middleware.CheckAuth(), middleware.UserAccess())
			sizes.GET("/", sizeHandler.GetAllSize)
		}

		categories := v1.Group("categories")
		{
			categories.Use(middleware.CheckAuth(), middleware.UserAccess())
			categories.GET("/", categoryHandler.GetAllCategory)
		}

		addOns := v1.Group("add-ons")
		{
			addOns.Use(middleware.CheckAuth(), middleware.UserAccess())
			addOns.GET("/", addOnHandler.GetAllAddOn)
		}
	}

	router.Run(cfg.Server.Url)
}
