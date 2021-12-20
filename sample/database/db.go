package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type NewDB struct {
	DB *gorm.DB
}

var dbInstance *NewDB

func GetDbConnection(dsn string) *NewDB {
	if dbInstance == nil {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err.Error())
		}
		dbInstance = &NewDB{db}
	}

	log.Printf("DB コネクション %p: ", dbInstance.DB)
	return dbInstance
}
