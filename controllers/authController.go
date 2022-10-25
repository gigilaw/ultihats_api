package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gigilaw/ultihats/config"
	"github.com/gigilaw/ultihats/handlers"
	"github.com/gigilaw/ultihats/initializers"
	"github.com/gigilaw/ultihats/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func UserEmailRegister(c *gin.Context) {
	var newUserBody struct {
		FirstName  string `binding:"required"`
		LastName   string `binding:"required"`
		Height     int    `binding:"required"`
		Gender     string `binding:"required"`
		Email      string `binding:"required,email"`
		Password   string `binding:"required,alphanum,min=8"`
		CommonName string `binding:"omitempty,alpha"`
		Birthday   string `binding:"required"`
	}
	if err := c.ShouldBind(&newUserBody); err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_VALIDATION["message"], err.Error()))
		return
	}
	passwordHash, err := models.HashPassword(newUserBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_HASH_PASSWORD["message"], config.ERROR_HASH_PASSWORD["details"]))
		return
	}

	birthday, err := models.ParseBirthday(newUserBody.Birthday)
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_PARSE_BIRTHDAY["message"], config.ERROR_PARSE_BIRTHDAY["details"]))
		return
	}

	user := models.User{
		FirstName:  newUserBody.FirstName,
		LastName:   newUserBody.LastName,
		Height:     newUserBody.Height,
		Gender:     newUserBody.Gender,
		Email:      newUserBody.Email,
		Birthday:   birthday,
		CommonName: newUserBody.CommonName,
		Password:   string(passwordHash),
	}

	if result := initializers.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_DATABASE["message"], result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, &user)
}

func UserLogin(c *gin.Context) {
	var login struct {
		Email    string `binding:"required"`
		Password string `binding:"required"`
	}

	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_VALIDATION["message"], err.Error()))
		return
	}
	var user models.User
	initializers.DB.First(&user, "email = ?", login.Email)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil || user.ID == 0 {
		c.JSON(http.StatusUnauthorized, handlers.ErrorMessage(config.ERROR_INVALID_LOGIN["message"], config.ERROR_INVALID_LOGIN["details"]))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage("ERROR_CREATE_TOKEN", "Failed to create token"))
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusBadRequest, gin.H{
		"login": "Login Success",
	})
}
