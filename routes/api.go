package routes

import (
	"net/http"

	"github.com/gigilaw/ultihats/controllers"
	"github.com/gigilaw/ultihats/middleware"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to UltiHats! Work in progress~",
		})
	})

	router.POST("/register/email", controllers.UserEmailRegister)
	router.POST("/login", controllers.UserLogin)

	// authorized := router.Group("/").Use(middleware.JWTAuth)
	// {

	// }
}