package main

import (
	"github.com/SSPinheiro/KalCaloria/config"
	"github.com/SSPinheiro/KalCaloria/router"
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

	router.Inicializar()
}
