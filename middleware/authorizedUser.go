package middleware

import (
	"net/http"
	"strconv"

	"github.com/gigilaw/ultihats/config"
	"github.com/gigilaw/ultihats/handlers"
	"github.com/gigilaw/ultihats/models"
	"github.com/gin-gonic/gin"
)

func AuthorizedUser(c *gin.Context) {
	user, _ := c.Get("user")
	userID, _ := strconv.ParseUint(c.Param("userID"), 0, 64)

	if user.(models.User).ID != uint(userID) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, handlers.ErrorMessage(config.ERROR_UNAUTHORIZED["message"], config.ERROR_UNAUTHORIZED["details"]))
	}
	c.Next()
}
