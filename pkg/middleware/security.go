package middleware

import (
	"os"

	"github.com/IvanTarjan/final-go-g5/pkg/web"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader("token")
		token := os.Getenv("TOKEN")

		if tokenHeader == "" || tokenHeader != token {
			web.Failure(ctx, 401, "Token inválido")
			return
		} else {
			ctx.Next()
		}
	}

}
