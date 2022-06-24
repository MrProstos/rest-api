package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type router struct {
	Router *mux.Router
}

func (router *router) Init() *router {
	router.Router = mux.NewRouter().StrictSlash(true)
	router.Router.HandleFunc("/registration/", Registration).Methods("GET")
	router.Router.HandleFunc("/auth/", Auth).Methods("GET")

	router.Router.HandleFunc("/client/{phone_num}", http.HandlerFunc(ShowClients)).Methods("GET")
	router.Router.HandleFunc("/client/add/", AuthMiddleware(http.HandlerFunc(AddClient))).Methods("POST")
	router.Router.HandleFunc("/client/update/", AuthMiddleware(http.HandlerFunc(UpdateClient))).Methods("POST")
	router.Router.HandleFunc("/client/delete/", AuthMiddleware(http.HandlerFunc(DelClient))).Methods("DELETE")

	router.Router.HandleFunc("/order/{phone_num}", AuthMiddleware(http.HandlerFunc(ShowOrder))).Methods("GET")
	router.Router.HandleFunc("/order/add/", AuthMiddleware(http.HandlerFunc(AddOrder))).Methods("POST")
	router.Router.HandleFunc("/order/update/", AuthMiddleware(http.HandlerFunc(UpdateOrder))).Methods("POST")
	router.Router.HandleFunc("/order/delete/", AuthMiddleware(http.HandlerFunc(DelOrder))).Methods("DELETE")
	return router
}

func NewRouter() *router {
	return &router{}
}
