package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type dataBaseConfig struct {
	dbName string
	dbPass string
	dbUser string
	dbHost string
}

func (database *dataBaseConfig) Connect() error {
	dbUrl := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		"localhost", "postgres", "postgres", "Zz123456")

	conn, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return err
	}
	db = conn
	return nil
}

func NewDataBaseConfig() *dataBaseConfig {
	return &dataBaseConfig{}
}

func GetConn() *gorm.DB {
	if db == nil {
		log.Fatalln("db not connected")
	}
	return db
}
