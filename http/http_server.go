package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MrProstos/rest-api/db"
	"github.com/gorilla/mux"
)

func getInfo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var client db.Client

	err = json.Unmarshal(msg, &client)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Add(db.DB); err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/client/", getInfo).Methods("POST")
	err := http.ListenAndServe(":2000", router)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server starting...")
}
