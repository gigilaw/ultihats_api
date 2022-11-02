package routes

import (
	"net/http"

	"github.com/gigilaw/ultihats/controllers"
	"github.com/gigilaw/ultihats/middleware"
	"github.com/gin-gonic/gin"
)

func ApiRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to UltiHats! Work in progress~",
		})
	})

	r.POST("/register/email", controllers.UserEmailRegister)
	r.POST("/login", controllers.UserLogin)

	authenticated := r.Group("/").Use(middleware.JWTAuth)
	{
		authenticated.GET("/user/:userID", controllers.GetUser)
		authenticated.POST("/user/:userID", middleware.AuthorizedUser, controllers.UpdateUser)

		authenticated.POST("/disc-skills", controllers.UpsertDiscSkills)
	}
	return r
}
