package db

import (
	"fmt"

	"github.com/MrProstos/rest-api/utils"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	db_name string = "postgres"
	db_pass string = "changeme"
	db_user string = "postgres"
	//db_type string = "postgres"
	db_host string = "localhost"
)

func init() {

	dbUrl := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", db_host, db_user, db_name, db_pass)

	conn, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		utils.Logger.Fatal(err.Error())
	}

	DB = conn
	DB.AutoMigrate(&Client{}, &Order{})
}

// возвращает дескриптор объекта DB
func GetDB() *gorm.DB {
	return DB
}
