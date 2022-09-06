package main

import (
	"log"
	"template/infra/datastore"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := datastore.NewDB()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	r := registry
}
