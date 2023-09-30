package handler

import (
	"log"
	"net/http"
	"online_wallet_humo/internal/service"
	"online_wallet_humo/pkg/models"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type CardHandler struct {
	CardService *service.CardService
}

func NewCardHandler(cardService *service.CardService) CardHandler {
	return CardHandler{cardService}
}

func (h *CardHandler) AddCardToUser(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	var card models.Card

	if err := c.BindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	cardNumberRegex := regexp.MustCompile(`^[0-9]+$`)
	if !cardNumberRegex.MatchString(card.Number) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Card number should contain only digits",
		})
		return
	}

	if len(card.Number) != 16 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Card should contain exectly 16 digits",
		})
		return
	}

	if !strings.HasPrefix(card.Number, "505827") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid beginning of card number",
		})
		return
	}

	card.UserID = uint(id)

	if err := h.CardService.AddCardToUser(&card); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Card created and added to user successfully",
	})
}

func (h *CardHandler) GetUserCards(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	cards, err := h.CardService.GetUserCards(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, cards)
}
