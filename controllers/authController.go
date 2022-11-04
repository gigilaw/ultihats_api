package controllers

import (
	"net/http"

	"github.com/gigilaw/ultihats/config"
	"github.com/gigilaw/ultihats/handlers"
	"github.com/gigilaw/ultihats/initializers"
	"github.com/gigilaw/ultihats/models"
	"github.com/gigilaw/ultihats/validation"
	"github.com/gin-gonic/gin"
)

func UserEmailRegister(c *gin.Context) {
	newUser := validation.NewUserBody
	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_VALIDATION["message"], err.Error()))
		return
	}

	passwordHash, err := handlers.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_HASH_PASSWORD["message"], config.ERROR_HASH_PASSWORD["details"]))
		return
	}

	birthday, err := models.ParseBirthday(newUser.Birthday)
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_PARSE_BIRTHDAY["message"], config.ERROR_PARSE_BIRTHDAY["details"]))
		return
	}

	user := models.User{
		FirstName:  newUser.FirstName,
		LastName:   newUser.LastName,
		Height:     newUser.Height,
		Gender:     newUser.Gender,
		Email:      newUser.Email,
		Birthday:   birthday,
		CommonName: newUser.CommonName,
		Password:   passwordHash,
	}

	if result := initializers.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_DATABASE["message"], result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, &user)
}

func UserLogin(c *gin.Context) {
	login := validation.Login
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_VALIDATION["message"], err.Error()))
		return
	}

	var user models.User
	initializers.DB.Preload("DiscSkills").First(&user, "email = ?", login.Email)

	if !handlers.JwtToken(c, config.Roles["USER"], &user.Password, &login.Password, &user.ID) {
		return
	}

	c.JSON(http.StatusOK, &user)
}

func OrganizationRegister(c *gin.Context) {
	newOrganization := validation.NewOrganizationBody

	if err := c.ShouldBind(&newOrganization); err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_VALIDATION["message"], err.Error()))
		return
	}

	passwordHash, err := handlers.HashPassword(newOrganization.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_HASH_PASSWORD["message"], config.ERROR_HASH_PASSWORD["details"]))
		return
	}

	organization := models.Organization{
		Name:           newOrganization.Name,
		Email:          newOrganization.Email,
		Password:       passwordHash,
		City:           newOrganization.City,
		Est:            newOrganization.Est,
		Facebook:       newOrganization.Facebook,
		Instagram:      newOrganization.Instagram,
		DisplayPicture: newOrganization.DisplayPicture,
	}

	if result := initializers.DB.Create(&organization); result.Error != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_DATABASE["message"], result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, &organization)
}

func OrganizationLogin(c *gin.Context) {
	login := validation.Login
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_VALIDATION["message"], err.Error()))
		return
	}

	var organization models.Organization
	initializers.DB.First(&organization, "email = ?", login.Email)

	if !handlers.JwtToken(c, config.Roles["ORGANIZATION"], &organization.Password, &login.Password, &organization.ID) {
		return
	}

	c.JSON(http.StatusOK, &organization)
}
