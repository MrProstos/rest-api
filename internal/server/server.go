package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

var Router *mux.Router

func init() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/registration/", Registration).Methods("GET")

	//router.HandleFunc("/", middleware(http.HandlerFunc(server.Auth))).Methods("GET")
	router.HandleFunc("/auth/", Auth).Methods("GET")
	//	router.HandleFunc("/db/",server.AuthMiddleware(http.HandlerFunc(server.)))

	router.HandleFunc("/showclietns/{phone_num}", AuthMiddleware(http.HandlerFunc(ShowClients))).Methods("GET")
	router.HandleFunc("/addclient/", AuthMiddleware(http.HandlerFunc(AddClient))).Methods("POST")
	router.HandleFunc("/updateclient/", AuthMiddleware(http.HandlerFunc(UpdateClient))).Methods("PUT")
	router.HandleFunc("/delclient/", AuthMiddleware(http.HandlerFunc(DelClient))).Methods("DELETE")

	router.HandleFunc("/showorder/{client_id}", AuthMiddleware(http.HandlerFunc(ShowOrder))).Methods("GET")
	router.HandleFunc("/addorder/", AuthMiddleware(http.HandlerFunc(AddOrder))).Methods("POST")
	router.HandleFunc("/updateorder/", AuthMiddleware(http.HandlerFunc(UpdateOrder))).Methods("PUT")
	router.HandleFunc("/delorder/", AuthMiddleware(http.HandlerFunc(DelOrder))).Methods("DELETE")
	Router = router
}

func GetRouter() *mux.Router {
	return Router
}
