package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gigilaw/ultihats/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func JwtToken(c *gin.Context, role string, savedPassword *string, loginPassword *string, id *uint) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(*savedPassword), []byte(*loginPassword)); err != nil || *id == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessage(config.ERROR_INVALID_LOGIN["message"], config.ERROR_INVALID_LOGIN["details"]))
		return false
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  id,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage("ERROR_CREATE_TOKEN", "Failed to create token"))
		return false
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	return true
}
