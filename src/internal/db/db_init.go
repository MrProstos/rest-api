package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	dbName string = "postgres"
	dbPass string = "changeme"
	dbUser string = "postgres"
	dbHost string = "localhost"
)

func init() {

	dbUrl := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPass)

	conn, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = conn
}

// GetDB возвращает дескриптор объекта DB
func GetDB() *gorm.DB {
	return DB
}
