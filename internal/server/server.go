package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

var Router *mux.Router

func init() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/registration/", Registration).Methods("GET")
	router.HandleFunc("/auth/", Auth).Methods("GET")

	router.HandleFunc("/client/{phone_num}", AuthMiddleware(http.HandlerFunc(ShowClients))).Methods("GET")
	router.HandleFunc("/client/add/", AuthMiddleware(http.HandlerFunc(AddClient))).Methods("POST")
	router.HandleFunc("/client/update/", AuthMiddleware(http.HandlerFunc(UpdateClient))).Methods("POST")
	router.HandleFunc("/client/delete/", AuthMiddleware(http.HandlerFunc(DelClient))).Methods("DELETE")

	router.HandleFunc("/order/{phone_num}", AuthMiddleware(http.HandlerFunc(ShowOrder))).Methods("GET")
	router.HandleFunc("/order/add/", AuthMiddleware(http.HandlerFunc(AddOrder))).Methods("POST")
	router.HandleFunc("/order/update/", AuthMiddleware(http.HandlerFunc(UpdateOrder))).Methods("POST")
	router.HandleFunc("/order/delete/", AuthMiddleware(http.HandlerFunc(DelOrder))).Methods("DELETE")
	Router = router

}

func GetRouter() *mux.Router {
	return Router
}
