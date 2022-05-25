package main

import (
	"log"
	"net/http"

	"github.com/MrProstos/rest-api/server"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/registration/", server.Registration).Methods("GET")

	//router.HandleFunc("/", middleware(http.HandlerFunc(server.Auth))).Methods("GET")
	router.HandleFunc("/auth/", server.Auth).Methods("GET")
	//	router.HandleFunc("/db/",server.Middleware(http.HandlerFunc(server.)))
	router.HandleFunc("/db/showclietns/{phone_num}", server.ShowClients).Methods("GET")
	router.HandleFunc("/db/addclient/", server.AddClient).Methods("POST")
	router.HandleFunc("/db/updateclient/", server.UpdateClient).Methods("PUT")
	router.HandleFunc("/db/delclient/", server.DelClient).Methods("DELETE")

	router.HandleFunc("/db/showorder/{client_id}", server.ShowOrder).Methods("GET")
	router.HandleFunc("/db/addorder/", server.AddOrder).Methods("POST")
	router.HandleFunc("/db/updateorder/", server.UpdateOrder).Methods("PUT")
	router.HandleFunc("/db/delorder/", server.DelOrder).Methods("DELETE")
	err := http.ListenAndServe(":2000", router)
	if err != nil {
		log.Fatal(err)
	}

}
