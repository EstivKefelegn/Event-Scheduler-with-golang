package middleware

import (
	"net/http"

	"Eventplanning.go/Api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User doesnt have authorization"})
		return
	}

	userID, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User doesnt have authorization"})
		return
	}
	context.Set("userID", userID)
	context.Next()
}
