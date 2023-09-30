package handler

import (
	"fmt"
	"log"
	"net/http"
	"online_wallet_humo/configs"
	"online_wallet_humo/internal/db"
	"online_wallet_humo/internal/repository"
	"online_wallet_humo/internal/service"
	"online_wallet_humo/pkg/models"

	"github.com/gin-gonic/gin"
)

var PostgresLine string

func InitRoutes() error {
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

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.SignUp)
			auth.POST("/login", authHandler.SignIn)
		}

		user := api.Group("/user", authHandler.UserIdentity)
		{
			user.POST("/", userHandler.CreateUser)
			user.GET("/:id", userHandler.GetUserByID)
			user.PUT("/", userHandler.UpdateUser)
			user.DELETE("/:id", userHandler.DeleteUser)
			user.POST("/avatar", userHandler.UploadAvatar)
			user.GET("/avatar", userHandler.GetAvatar)
		}

		service := api.Group("/service", authHandler.UserIdentity)
		{
			service.POST("/", serviceHandler.CreateService)
			service.GET("/:id", serviceHandler.GetServiceByID)
			service.PUT("/:id", serviceHandler.UpdateServiceByID)
			service.DELETE("/:id", serviceHandler.DeleteServiceByID)
		}

		card := api.Group("/card", authHandler.UserIdentity)
		{
			card.POST("/", cardHandler.AddCardToUser)
			card.GET("/", cardHandler.GetUserCards)
		}

		transfer := api.Group("/transfer", authHandler.UserIdentity)
		{
			transfer.POST("/from_wallet_to_wallet", transferHandler.TransferFromWalletToWallet)
			transfer.POST("/from_card_to_wallet", transferHandler.TransferFromCardToWallet)
			transfer.GET("/transactions", transferHandler.GetUserTransactions)
		}

		favorite := api.Group("/favorite", authHandler.UserIdentity)
		{
			favorite.POST("/", favHandler.CreateFavorite)
		}

		stats := api.Group("/statistics", authHandler.UserIdentity)
		{
			stats.POST("/", staHandler.GetStatistics)
		}
	}

	config, err := configs.InitConfigs()
	if err != nil {
		log.Println(err)
		return err
	}

	address := config.Server.Host + config.Server.Port

	srv := http.Server{
		Addr:    address,
		Handler: r,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
		return err
	}

	PostgresLine = ToStringDBConfig(config)

	return nil
}

func ToStringDBConfig(c *models.Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.Database.Host, c.Database.User, c.Database.Password, c.Database.BDName, c.Database.Port, c.Database.SSLMode,
	)
}
