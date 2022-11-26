package router

import (
	"final-project-backend/handler"
	"final-project-backend/middleware"
	"final-project-backend/pkg/helper"
	"final-project-backend/repository"
	"final-project-backend/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouterSetUp(router *gin.Engine, db *gorm.DB) {
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

	promoRepository := repository.NewPromoRepositoryImpl(db)
	promoUsecase := usecase.NewPromoUsecaseImpl(promoRepository)
	promoHandler := handler.NewPromoHandler(promoUsecase)

	promoUserRepository := repository.NewPromoUserRepositoryImpl(db)
	promoUserUsecase := usecase.NewPromoUserUsecaseImpl(promoUserRepository)
	promoUserHandler := handler.NewPromoUserHandler(promoUserUsecase)

	paymentRepository := repository.NewPaymentRepositoryImpl(db)
	paymentUsecase := usecase.NewPaymentUsecaseImpl(paymentRepository, userRepository, promoUserRepository)
	paymenthandler := handler.NewPaymentHandler(paymentUsecase)

	shippingRepository := repository.NewShippingRepositoryImpl(db)
	shippingUsecase := usecase.NewShippingUsecaseImpl(
		shippingRepository, paymentRepository, userRepository,
		addressRepository, sizeRepository, categoryRepository, addOnRepository,
	)
	shippingHandler := handler.NewShippingHandler(shippingUsecase)

	gameUsecase := usecase.NewGameUsecaseImpl(promoRepository, promoUserRepository, shippingRepository)
	gameHandler := handler.NewGameHandler(gameUsecase)

	router.NoRoute(func(c *gin.Context) {
		helper.ErrorResponse(c.Writer, "not found", 404)
	})
	router.Use(middleware.CORSMiddleware())

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

		shippings := v1.Group("shippings")
		{
			shippings.Use(middleware.CheckAuth(), middleware.AdminAccess())
			shippings.GET("/", shippingHandler.GetAllShipping)
			shippings.GET("/:id", shippingHandler.GetShippingById)
			shippings.PUT("/:id/status", shippingHandler.UpdateStatusShipping)
		}

		promos := v1.Group("promos")
		{
			promos.Use(middleware.CheckAuth(), middleware.AdminAccess())
			promos.GET("/", promoHandler.GetAllPromo)
			promos.POST("/", promoHandler.CreatePromo)
			promos.PUT("/:id", promoHandler.UpdatePromo)
		}

		users := v1.Group("users")
		{
			users.Use(middleware.CheckAuth(), middleware.UserAccess())
			users.GET("/", userHandler.GetUserById)
			users.PUT("/", userHandler.UpdateUserById)
			users.POST("/top-up", userHandler.TopUp)

			addresses := users.Group("addresses")
			{
				addresses.GET("/", addressHandler.GetAllAddressByUserId)
				addresses.GET("/:id", addressHandler.GetAddressByUserId)
				addresses.POST("/", addressHandler.CreateAddress)
				addresses.PUT("/:id", addressHandler.UpdateAddressByUserId)
			}

			shippings := users.Group("shippings")
			{
				shippings.GET("/", shippingHandler.GetAllShippingByUserId)
				shippings.GET("/:id", shippingHandler.GetShippingByUserId)
				shippings.POST("/", shippingHandler.CreateShipping)
				shippings.PUT("/:id/review", shippingHandler.UpdateReviewByUserId)

				payments := shippings.Group("payments")
				{
					payments.PUT("/:id", paymenthandler.PayUserShipping)
				}
			}

			games := users.Group("games")
			{
				games.POST("/play", gameHandler.Play)
			}

			promos := users.Group("promos")
			{
				promos.GET("/", promoUserHandler.GetAllPromoUserByUserId)
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
}
