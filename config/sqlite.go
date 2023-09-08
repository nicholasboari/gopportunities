package config

import (
	"os"

	"github.com/nicholasboari/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	pathDb := "./db/main.db"

	// check if the database file exists
	_, err := os.Stat(pathDb)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("error to create directory: %v", err)
			return nil, err
		}
		file, err := os.Create(pathDb)

		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// create db and connect
	db, err := gorm.Open(sqlite.Open(pathDb), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	// migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}
	return db, nil
}
