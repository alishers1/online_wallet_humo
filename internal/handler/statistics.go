package handler

import (
	"net/http"
	"online_wallet_humo/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type StartAndEndDate struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type StatisticsHandler struct {
	staService *service.StatisticsService
}

func NewStatisticsHandler(staService *service.StatisticsService) *StatisticsHandler {
	return &StatisticsHandler{staService}
}

func (h *StatisticsHandler) GetStatistics(c *gin.Context) {
	_, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		logrus.Error(err)
		return
	}

	var request StartAndEndDate

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logrus.Error(err)
		return
	}

	userCount, err := h.staService.GetUsersCount(request.StartDate, request.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		logrus.Error(err)
		return
	}

	cardCount, err := h.staService.GetCardsCount(request.StartDate, request.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logrus.Error(err)
		return
	}

	transactionCount, err := h.staService.GetTransactionsCount(request.StartDate, request.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logrus.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users_count":        userCount,
		"cards_count":        cardCount,
		"transactions_count": transactionCount,
	})
}
