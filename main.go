package main

import (
	"github.com/gigilaw/ultihats/initializers"
	"github.com/gigilaw/ultihats/routes"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := routes.ApiRoutes()
	r.Run()
}
