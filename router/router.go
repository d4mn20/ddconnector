package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize the Router using default Gin configuration
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	fmt.Println("\nRouter initialized... RUNNING...")
	// Listen and serve on 0.0.0.0:21777
	router.Run(":21777")
}
