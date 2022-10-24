package handlers

import "github.com/gin-gonic/gin"

func ErrorMessage(message string, details string) gin.H {
	err := map[string]string{"message": message, "details": details}

	return gin.H{
		"error": err,
	}
}
