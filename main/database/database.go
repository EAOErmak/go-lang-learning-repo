package database

import (
	"errors"
	"os"
	"strings"

	"go-learn/main/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := strings.TrimSpace(os.Getenv("DATABASE_URL"))
	if dsn == "" {
		return errors.New("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(
		&models.DictionaryItem{},
		&models.DiaryEntry{},
		&models.EntryMetric{},
		&models.EntryMetricValue{},
	); err != nil {
		return err
	}

	DB = db
	return nil
}
