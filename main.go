package main

import (
	"github.com/gigilaw/ultihats/initalizers"
	"github.com/gigilaw/ultihats/routes"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectDB()
}

func main() {
	router = gin.Default()
	routes.ApiRoutes(router)
	router.Run()
}
