package main

import (
	"github.com/gigilaw/ultihats/initializers"
	"github.com/gigilaw/ultihats/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.User{},
		&models.DiscSkills{},
		&models.Organization{},
	)
}
