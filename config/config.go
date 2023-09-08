package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return errors.New("Erro fake!")
	// return nil
}

func GetLogger(p string) *Logger {
	return NewLogger(p)
}
