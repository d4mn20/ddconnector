package main

import (
	"fmt"

	"dev.azure.com/bbts-lab/DevSecOps/_git/ms-ddconnector/config"
	"dev.azure.com/bbts-lab/DevSecOps/_git/ms-ddconnector/router"
)

func main() {
	logger := *config.GetLogger("main")
	// Initialize Config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	fmt.Println("Starting...")
	// Initialize Router
	router.Initialize()
}
