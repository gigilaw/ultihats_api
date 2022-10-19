package controllers

import (
	"net/http"

	"github.com/gigilaw/ultihats/initalizers"
	"github.com/gigilaw/ultihats/models"
	"github.com/gin-gonic/gin"
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

	if result := initalizers.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, &user)
}
