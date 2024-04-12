package controllers

import (
	"ABC_PHARMACY/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

type Repository struct {
	DB *gorm.DB
}

var r *Repository

func NewConnection(config *Config) error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = models.MigrateDB(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	r = &Repository{
		DB: db,
	}

	return nil

}
