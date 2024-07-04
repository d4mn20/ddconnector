package router

import (
	"fmt"

	"dev.azure.com/bbts-lab/DevSecOps/_git/ddconnector/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	fmt.Println("Initializing Routes...")
	handler.InitializeHandler()

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/health", handler.HealthHandler)
		v1.POST("/upload", handler.PublishScanHandler)
	}
}
