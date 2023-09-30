package handler

import (
	"log"
	"net/http"
	"online_wallet_humo/internal/service"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type signIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signUp struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var request signUp

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	if !strings.Contains(request.Email, "@") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong email format",
		})
		return
	}

	phoneNumberRegex := regexp.MustCompile(`^[0-9]+$`)
	if !phoneNumberRegex.MatchString(request.PhoneNumber) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Phone number should contain only digits",
		})
		return
	}

	id, err := h.authService.CreateUser(
		request.FullName, request.Email, request.PhoneNumber, request.Password,
	)
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

func (h *AuthHandler) SignIn(c *gin.Context) {
	var request signIn

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	if !strings.Contains(request.Email, "@") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email address",
		})
		return 
	}

	token, err := h.authService.GenerateToken(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println("error in auth.go/handler/line 75")
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
