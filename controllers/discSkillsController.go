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

func UpsertDiscSkills(c *gin.Context) {
	userValue, _ := c.Get("user")
	user := models.User(userValue.(models.User))
	userId := user.ID

	discSkillsValidation := validation.UpsertDiscSkills
	if err := c.ShouldBind(&discSkillsValidation); err != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_VALIDATION["message"], err.Error()))
		return
	}

	upsertDiscSkills := models.DiscSkills{
		PrimaryRole:         discSkillsValidation.PrimaryRole,
		Throwing:            discSkillsValidation.Throwing,
		Catching:            discSkillsValidation.Catching,
		OffensiveStrategies: discSkillsValidation.OffensiveStrategies,
		DefensiveStrategies: discSkillsValidation.DefensiveStrategies,
		Public:              discSkillsValidation.Public,
	}

	var discSkills models.DiscSkills
	result := initializers.DB.Where(models.DiscSkills{UserID: userId}).Assign(&upsertDiscSkills).FirstOrCreate(&discSkills)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, handlers.ErrorMessage(config.ERROR_DATABASE["message"], result.Error.Error()))
		return
	}

	user.DiscSkills = discSkills

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
