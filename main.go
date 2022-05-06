package main

import (
	"fmt"

	"github.com/MrProstos/rest-api/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", db.Db_host, db.Db_user, db.Db_name, db.Db_pass)
	fmt.Println(dbUri)

	conn, err := gorm.Open(db.Db_type, dbUri)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(conn.Error)
	user := db.Operator{Username: "Vlad", Token: "test"}
	user.Del(conn)

}
