package controllers

import (
	"net/http"
	"time"

	"github.com/gigilaw/ultihats/config"
	"github.com/gigilaw/ultihats/handlers"
	"github.com/gigilaw/ultihats/initializers"
	"github.com/gigilaw/ultihats/models"
	"github.com/gigilaw/ultihats/validation"
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetUser(c *gin.Context) {
	var user models.User
	initializers.DB.Preload("DiscSkills").First(&user, c.Param("userID"))

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("userID")

	updateUserBody := validation.UpdateUserBody
	if err := c.ShouldBind(&updateUserBody); err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_VALIDATION["message"], err.Error()))
		return
	}

	var birthday time.Time
	if len(updateUserBody.Birthday) > 0 {
		parsedBirthday, err := models.ParseBirthday(updateUserBody.Birthday)

		if err != nil {
			c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_PARSE_BIRTHDAY["message"], config.ERROR_PARSE_BIRTHDAY["details"]))
			return
		}
		birthday = parsedBirthday
	}

	var password string
	if len(updateUserBody.Password) > 0 {
		hashedPassword, err := models.HashPassword(updateUserBody.Password)

		if err != nil {
			c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_HASH_PASSWORD["message"], config.ERROR_HASH_PASSWORD["details"]))
			return
		}
		password = hashedPassword
	}

	updateUser := models.User{
		FirstName:      updateUserBody.FirstName,
		LastName:       updateUserBody.LastName,
		Height:         updateUserBody.Height,
		Gender:         updateUserBody.Gender,
		Email:          updateUserBody.Email,
		CommonName:     updateUserBody.CommonName,
		DisplayPicture: updateUserBody.DisplayPicture,
		Password:       password,
		Birthday:       birthday,
	}

	var user models.User
	initializers.DB.First(&user, userID)
	initializers.DB.Model(&user).Updates(updateUser)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
