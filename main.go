package main

import (
	"github.com/nicholasboari/gopportunities/config"
	"github.com/nicholasboari/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	router.Initialize()
}
