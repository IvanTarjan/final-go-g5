package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader("token")
		token := os.Getenv("TOKEN")

		if tokenHeader == "" || tokenHeader != token {
			// web.Failure(ctx, 401, "Token inválido")
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "Invalid token.",
			})
			return
		} else {
			ctx.Next()
		}
	}

}
