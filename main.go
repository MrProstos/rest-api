package main

import (
	"log"
	"net/http"

	"github.com/MrProstos/rest-api/server"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/addclient/", server.AddClient).Methods("POST")
	router.HandleFunc("/updateclient/", server.UpdateClient).Methods("POST")
	router.HandleFunc("/delclient/", server.DelClient).Methods("POST")
	router.HandleFunc("/addorder/", server.AddOrder).Methods("POST")
	router.HandleFunc("/updateorder/", server.UpdateOrder).Methods("POST")
	router.HandleFunc("/delorder/", server.DelOrder).Methods("POST")
	err := http.ListenAndServe(":2000", router)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server starting...")
}
