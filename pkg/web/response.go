package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Response(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, data)
}

type response struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

// Success escribe una respuesta exitosa
func Success(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, response{
		Data:   data,
		Status: status,
	})
}

func Failure(ctx *gin.Context, status int, format string, args ...interface{}) {
	err := errorResponse{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}
	Response(ctx, status, err)
}
