package main

import (
	"github.com/MrProstos/rest-api/internal/gateway/db"
	"github.com/MrProstos/rest-api/internal/gateway/myldap"
	"github.com/MrProstos/rest-api/internal/server"
	"log"
	"net/http"
)

func main() {
	err := db.NewDataBaseConfig().Connect()
	if err != nil {
		log.Fatalln(err)
	}
	err = myldap.NewConf(myldap.Url, myldap.Bind, myldap.Password).Connect()
	if err != nil {
		log.Fatalln(err)
	}
	err = http.ListenAndServe(":2000", server.NewRouter().Init())
	if err != nil {
		log.Fatal(err)
	}
}
