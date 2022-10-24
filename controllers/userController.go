package controllers

import (
	"net/http"
	"time"

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
	userID := c.Param("userID")

	var updateUserBody struct {
		FirstName  string `binding:"omitempty,alpha"`
		LastName   string `binding:"omitempty,alpha"`
		Height     int    `binding:"omitempty,numeric"`
		Gender     string `binding:"omitempty,alpha"`
		Email      string `binding:"omitempty,email"`
		Password   string `binding:"omitempty,alphanum,min=8"`
		CommonName string `binding:"omitempty,alpha"`
		Birthday   string
	}
	if err := c.ShouldBind(&updateUserBody); err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_VALIDATION["message"], err.Error()))
		return
	}

	var birthday time.Time
	if len(updateUserBody.Birthday) > 0 {
		parsedBirthday, err := models.ParseBirthday(updateUserBody.Birthday)

		if err != nil {
			c.JSON(http.StatusBadRequest, handlers.ErrorMessage("ERROR_PARSE_BIRTHDAY", "Failed to parse birthday"))
			return
		}
		birthday = parsedBirthday
	}

	var password string
	if len(updateUserBody.Password) > 0 {
		hashedPassword, err := models.HashPassword(updateUserBody.Password)

		if err != nil {
			c.JSON(http.StatusBadRequest, handlers.ErrorMessage("ERROR_PARSE_BIRTHDAY", "Failed to parse birthday"))
			return
		}
		password = hashedPassword
	}

	updateUser := models.User{
		FirstName:  updateUserBody.FirstName,
		LastName:   updateUserBody.LastName,
		Height:     updateUserBody.Height,
		Gender:     updateUserBody.Gender,
		Email:      updateUserBody.Email,
		Birthday:   birthday,
		CommonName: updateUserBody.CommonName,
		Password:   password,
	}

	var user models.User
	initializers.DB.First(&user, userID)
	initializers.DB.Model(&user).Updates(updateUser)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
