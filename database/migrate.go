package main

import (
	"github.com/gigilaw/ultihats/initalizers"
	"github.com/gigilaw/ultihats/models"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectDB()
}

func main() {
	initalizers.DB.AutoMigrate(&models.User{})
}
