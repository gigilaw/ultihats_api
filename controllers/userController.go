package controllers

import (
	"net/http"
	"strconv"

	"github.com/gigilaw/ultihats/config"
	"github.com/gigilaw/ultihats/handlers"
	"github.com/gigilaw/ultihats/initializers"
	"github.com/gigilaw/ultihats/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID := c.Param("userID")
	var user models.User

	initializers.DB.First(&user, userID)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UpdateUser(c *gin.Context) {
	user, _ := c.Get("user")
	userID, _ := strconv.ParseUint(c.Param("userID"), 0, 64)

	if user.(models.User).ID != uint(userID) {
		handlers.Error(c, http.StatusUnauthorized, config.ERROR_UNAUTHORIZED["message"], config.ERROR_UNAUTHORIZED["details"])
	}
}
