package handler

import (
	"github.com/SSPinheiro/KalCaloria/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializarHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
