package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}
	return nil
}

func GetLogger(p string) *Logger {

	logger = newLogger(p)
	return logger
}
