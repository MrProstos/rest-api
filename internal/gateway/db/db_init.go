package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

const (
	dbName string = "postgres"
	dbPass string = "Zz123456"
	dbUser string = "postgres"
	dbHost string = "localhost"
)

func init() {
	DB = connPSQL()
}

func connPSQL() *gorm.DB {
	dbUrl := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPass)

	conn, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = conn.AutoMigrate(&Client{}, &Order{})
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

// GetDB возвращает дескриптор объекта DB
func GetDB() *gorm.DB {
	return DB
}
