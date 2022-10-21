package handlers

import "github.com/gin-gonic/gin"

func Error(c *gin.Context, httpStatus int, message string, details string) {
	err := map[string]string{"message": message, "details": details}

	c.JSON(httpStatus, gin.H{
		"error": err,
	})
}
