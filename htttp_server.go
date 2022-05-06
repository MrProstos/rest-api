package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MrProstos/rest-api/db"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Operator_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		user := db.Operator{}
		json.Unmarshal(body, &user)
		user.Token = uuid.New().String()

		fmt.Fprintf(w, "%v %v", user.Username, user.Token)
	}
}

var (
	db_name string = "Users"
	db_pass string = "changeme"
	db_user string = "postgres"
	db_type string = "postgres"
	db_host string = "localhost"
	//db_port int    = 5434
)

func main() {
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", db_host, db_user, db_name, db_pass)
	fmt.Println(dbUri)

	conn, err := gorm.Open(db_type, dbUri)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(conn.Error)
	user := db.Operator{Username: "Vlad", Token: "test"}
	user.Del(conn)

}
