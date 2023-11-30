package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Description do ping
func main() {
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()

	// Create a new Gin engine.
	router := gin.New()

	// Add the swagger handler.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
