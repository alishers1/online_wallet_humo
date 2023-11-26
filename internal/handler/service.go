package handler

import (
	"log"
	"net/http"
	"online_wallet_humo/internal/service"
	"online_wallet_humo/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	serviceService *service.ServiceService
}

func NewServiceHandler(s *service.ServiceService) *ServiceHandler {
	return &ServiceHandler{s}
}

func (h *ServiceHandler) CreateService(c *gin.Context) {
	var service models.Service

	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	if err := c.BindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	user, err := h.serviceService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	if !user.IsAdmin {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "you are not admin",
		})
		return
	}

	id, err := h.serviceService.CreateService(&service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *ServiceHandler) GetServiceByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	user, err := h.serviceService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	if !user.IsAdmin {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "you are not admin",
		})
		return
	}

	service, err := h.serviceService.GetServiceByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, service)
}

func (h *ServiceHandler) UpdateServiceByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	user, err := h.serviceService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	if !user.IsAdmin {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "you are not admin",
		})
		return
	}

	var service models.Service
	if err := c.BindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	service.ID = uint(id)

	if err := h.serviceService.UpdateServiceByID(&service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Service successfully updated",
	})
}

func (h *ServiceHandler) DeleteServiceByID(c *gin.Context) {
	idStr := c.Param("id")
	serviceID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	user, err := h.serviceService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	if !user.IsAdmin {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "you are not admin",
		})
		return
	}

	if err := h.serviceService.DeleteServiceByID(uint(serviceID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Service successfully deleted",
	})
}
