package datastore

import (
	"fmt"
	"log"
	"template/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	config, _ := config.GetConfig()
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%t",
		config.Host, config.Port, config.User, config.DBPassword, config.DBName, config.SSLMode)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
