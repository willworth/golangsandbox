package middlewares

import (
	"net/http"

	"example.com/example/10restAPI/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No JWT bearer token detected"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Not authorized"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
