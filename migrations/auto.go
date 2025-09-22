package main

import (
	"log"
	"os"

	"go-adv-demo/internal/link"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatalln("failed to read DSN: DSN is not provided in ENV")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")))
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&link.Link{})
	if err != nil {
		panic(err)
	}
}
