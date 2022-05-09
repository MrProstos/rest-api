package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

var (
	db_name string = "postgres"
	db_pass string = "changeme"
	db_user string = "postgres"
	db_type string = "postgres"
	db_host string = "localhost"
)

func init() {

	dbUrl := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", db_host, db_user, db_name, db_pass)

	conn, err := gorm.Open(db_type, dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	DB = conn

}
