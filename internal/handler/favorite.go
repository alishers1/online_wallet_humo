package handler

import (
	"log"
	"net/http"
	"online_wallet_humo/internal/service"
	"online_wallet_humo/pkg/models"

	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct {
	favService *service.FavoriteService
}

func NewFavoriteHandler(favService *service.FavoriteService) *FavoriteHandler {
	return &FavoriteHandler{favService}
}

func (h *FavoriteHandler) CreateFavorite(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	var favorite models.Favorite

	if err := c.BindJSON(&favorite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	favorite.UserID = id

	if err := h.favService.CreateFavorite(favorite); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction added to favorites",
	})
}