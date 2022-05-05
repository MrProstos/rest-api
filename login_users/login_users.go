package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type LoginUser interface {
	Login(User) error
}

type User struct {
	Name     string
	Password string
}

func (user User) Login(User) error {
	fmt.Println(user.Name, user.Password)
	return nil
}

var db *gorm.DB

func main() {
	e := godotenv.Load("/root/go/src/github.com/MrProstos/rest-api/login_users/.env")
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(conn)
}
