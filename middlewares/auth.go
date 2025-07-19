package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zin-min-thu/go-rest-api/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	context.Set("userId", userId)

	// next request handler will be excute
	context.Next()
}
