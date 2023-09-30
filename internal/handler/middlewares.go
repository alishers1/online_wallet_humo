package handler

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

const userCtx = "UserID"

func (h *AuthHandler) UserIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": errors.New("Empty auth header"),
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"reason": "token is empty",
		})
		return
	}

	userID, err := h.authService.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	c.Set(userCtx, userID)
}

func getUserID(c *gin.Context) (uint, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user id not found",
		})
		return 0, errors.New("user not found")
	}

	return cast.ToUint(id), nil
}
