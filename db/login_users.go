package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUrl := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", username, password, dbName, dbHost)
	fmt.Println(dbUrl)

	conn, err := gorm.Open("postgres", dbUrl)
	if err != nil {
		fmt.Print(err)
	}
	DB = conn
	fmt.Println(conn)
}
