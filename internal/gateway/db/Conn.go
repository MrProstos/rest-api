package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type dataBaseConfig struct {
	dbName string
	dbPass string
	dbUser string
	dbHost string
}

func (database *dataBaseConfig) SetConnect(dbName string, dbPass string, dbUser string, dbHost string) *dataBaseConfig {
	database.dbName = dbName
	database.dbPass = dbPass
	database.dbUser = dbUser
	database.dbHost = dbHost
	return database
}

func (database *dataBaseConfig) Connect() error {
	dbUrl := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		database.dbHost, database.dbUser, database.dbName, database.dbPass)

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
	return db
}
