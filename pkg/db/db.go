package db

import (
	"log"

	"go-adv-demo/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	if conf.Db.Dsn == "" {
		log.Fatalln("failed to read DSN: DSN is not provided in ENV")
	}

	db, err := gorm.Open(postgres.Open(conf.Db.Dsn))
	if err != nil {
		panic(err)
	}

	return &Db{db}
}
