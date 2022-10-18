package main

import (
	"github.com/gigilaw/ultihats/config"
	"github.com/gigilaw/ultihats/initalizers"
	"github.com/gin-gonic/gin"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(constants.HTTP_SUCCESS, gin.H{
			"message": "Welcome to UltiHats! Work in progress~",
		})
	})
	r.Run()
}
