// @title Gopportunities API
// @version 1.0
// @description API for managing job openings
// @BasePath /api/v1
package main

import (
	"github.com/ettory-automation/gopportunities/config"
	r "github.com/ettory-automation/gopportunities/router"
	_ "github.com/ettory-automation/gopportunities/docs"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	r.Initialize()
}
