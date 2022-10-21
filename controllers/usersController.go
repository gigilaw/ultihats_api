package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gigilaw/ultihats/initializers"
	"github.com/gigilaw/ultihats/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func UserEmailRegister(c *gin.Context) {
	var newUser struct {
		FirstName  string `binding:"required"`
		LastName   string `binding:"required"`
		Height     int    `binding:"required"`
		Gender     string `binding:"required"`
		Email      string `binding:"required,email"`
		Password   string `binding:"required,alphanum,min=8"`
		Birthday   string `binding:"required"`
		CommonName string
	}
	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.User{
		FirstName:  newUser.FirstName,
		LastName:   newUser.LastName,
		Height:     newUser.Height,
		Gender:     newUser.Gender,
		Email:      newUser.Email,
		Birthday:   newUser.Birthday,
		CommonName: newUser.CommonName,
		Password:   string(hash),
	}

	if result := initializers.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}
	var user models.User
	initializers.DB.First(&user, "email = ?", login.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusBadRequest, gin.H{
		"login": "Login Success",
	})
}
