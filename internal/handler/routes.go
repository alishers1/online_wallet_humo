package handler

import (
	"log"
	"net/http"
	"online_wallet_humo/internal/db"
	"online_wallet_humo/internal/repository"
	"online_wallet_humo/internal/service"

	"github.com/gin-gonic/gin"
)

var PostgresLine string

func InitRoutes(serverStr string) error {
	r := gin.Default()
	db := db.GetDBConn()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	transferRepo := repository.NewTransactionRepository(db)
	transferService := service.NewTransferService(transferRepo)
	transferHandler := NewTransferHandler(transferService)

	serviceRepo := repository.NewServiceRepository(db)
	serviceService := service.NewServiceService(serviceRepo)
	serviceHandler := NewServiceHandler(serviceService)

	authRepo := repository.NewAuthPostgres(db)
	authService := service.NewAuthService(*authRepo)
	authHandler := NewAuthHandler(*authService)

	cardRepo := repository.NewCardRepository(db)
	cardService := service.NewCardService(cardRepo)
	cardHandler := NewCardHandler(cardService)

	favRepo := repository.NewFavoriteRepository(db)
	favService := service.NewFavoriteService(favRepo)
	favHandler := NewFavoriteHandler(favService)

	staRepo := repository.NewStatisticsRepository(db)
	staService := service.NewStatisticsService(staRepo)
	staHandler := NewStatisticsHandler(staService)

	apiUser := r.Group("/api/v1")
	{
		auth := apiUser.Group("/auth")
		{
			auth.POST("/register", userHandler.SignUpUser)
			auth.POST("/login", userHandler.SignInUser)
		}

		user := apiUser.Group("/user", authHandler.UserIdentity)
		{
			user.POST("/avatar", userHandler.UploadAvatarByUser)
			user.GET("/avatar", userHandler.GetAvatarByUser)
		}

		service := apiUser.Group("/service", authHandler.UserIdentity)
		{
			service.GET("/:id", serviceHandler.GetServiceByID)
		}

		transaction := apiUser.Group("/transaction", authHandler.UserIdentity)
		{
			transaction.POST("/from_wallet_to_wallet", transferHandler.TransferFromWalletToWallet)
			transaction.POST("/from_card_to_wallet", transferHandler.TransferFromCardToWallet)
			transaction.GET("/", transferHandler.GetUserTransactions)
			transaction.POST("/add_favorite", favHandler.CreateFavorite)
		}

		card := apiUser.Group("/card", authHandler.UserIdentity)
		{
			card.POST("/", cardHandler.AddCardToUser)
			card.GET("/", cardHandler.GetUserCards)
		}
	}

	apiAdmin := r.Group("/api/v1/admin")
	{
		auth := apiAdmin.Group("/auth")
		{
			auth.POST("/register", authHandler.SignUp)
			auth.POST("/login", authHandler.SignIn)
		}

		user := apiAdmin.Group("/user", authHandler.UserIdentity)
		{
			user.POST("/", userHandler.CreateUserByAdmin)
			user.GET("/:id", userHandler.GetUserByID)
			user.PUT("/:id", userHandler.UpdateUserByAdmin)
			user.DELETE("/:id", userHandler.DeleteUser)
			user.POST("/avatar/:id", userHandler.UploadAvatarByAdmin)
			user.GET("/avatar/:id", userHandler.GetAvatarByAdmin)
		}

		service := apiAdmin.Group("/service", authHandler.UserIdentity)
		{
			service.POST("/", serviceHandler.CreateService)
			service.GET("/:id", serviceHandler.GetServiceByID)
			service.PUT("/:id", serviceHandler.UpdateServiceByID)
			service.DELETE("/:id", serviceHandler.DeleteServiceByID)
		}

		stats := apiAdmin.Group("/statistics", authHandler.UserIdentity)
		{
			stats.POST("/", staHandler.GetStatistics)
		}

		transfer := apiAdmin.Group("/transfer", authHandler.UserIdentity)
		{
			transfer.GET("/transactions/:id", transferHandler.GetUserTransactionsByAdmin)
		}
	}

	srv := http.Server{
		Addr:    serverStr,
		Handler: r,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
