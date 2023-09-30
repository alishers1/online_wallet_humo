package handler

import (
	"log"
	"net/http"
	"online_wallet_humo/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TransferMoney struct {
	RecipientID uint    `json:"recipient_id"`
	CardID      uint    `json:"card_id"`
	ServiceID   uint    `json:"service_id"`
	Amount      float64 `json:"amount"`
}

type TransferHandler struct {
	transferService *service.TransferService
}

func NewTransferHandler(ts *service.TransferService) *TransferHandler {
	return &TransferHandler{ts}
}

func (h *TransferHandler) TransferFromWalletToWallet(c *gin.Context) {
	var tm TransferMoney

	senderID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	if err := c.BindJSON(&tm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	if err := h.transferService.TransferFromWalletToWallet(senderID, tm.RecipientID, tm.ServiceID, tm.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction created successfully",
	})
}

func (h *TransferHandler) TransferFromCardToWallet(c *gin.Context) {
	var tm TransferMoney

	senderID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		logrus.Error(err)
		return
	}

	if err := c.BindJSON(&tm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logrus.Error(err)
		return
	}

	if err := h.transferService.TransferFromCardToWallet(senderID, tm.RecipientID, tm.CardID, tm.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		logrus.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction created successfully",
	})
}

func (h *TransferHandler) GetUserTransactions(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	transactions, err := h.transferService.GetUserTransactions(userID, uint(page), uint(pageSize))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, transactions)
}
