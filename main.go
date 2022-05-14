package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MrProstos/rest-api/server"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/addclient/", server.EnterClient).Methods("POST")
	err := http.ListenAndServe(":2000", router)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server starting...")
}
