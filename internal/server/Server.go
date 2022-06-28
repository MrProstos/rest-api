package server

import (
	"github.com/MrProstos/rest-api/internal/gateway/db"
	"github.com/MrProstos/rest-api/internal/gateway/myldap"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	Router *mux.Router
}

func (router *Server) Init() *mux.Router {
	err := db.NewDataBaseConfig().Connect()
	if err != nil {
		log.Fatalln(err)
	}
	err = myldap.NewConf(myldap.Url, myldap.Bind, myldap.Password).Connect()
	if err != nil {
		log.Fatalln(err)
	}

	router.Router = mux.NewRouter().StrictSlash(true)
	router.Router.HandleFunc("/registration/", Registration).Methods("POST")
	router.Router.HandleFunc("/auth/", Auth).Methods("POST")

	router.Router.HandleFunc("/client/{phone_num}", AuthMiddleware(http.HandlerFunc(ShowClients))).Methods("GET")
	router.Router.HandleFunc("/client/add/", AuthMiddleware(http.HandlerFunc(AddClient))).Methods("POST")
	router.Router.HandleFunc("/client/update/", AuthMiddleware(http.HandlerFunc(UpdateClient))).Methods("PUT")
	router.Router.HandleFunc("/client/delete/", AuthMiddleware(http.HandlerFunc(DelClient))).Methods("DELETE")

	router.Router.HandleFunc("/order/{phone_num}", AuthMiddleware(http.HandlerFunc(ShowOrder))).Methods("GET")
	router.Router.HandleFunc("/order/add/", AuthMiddleware(http.HandlerFunc(AddOrder))).Methods("POST")
	router.Router.HandleFunc("/order/update/", AuthMiddleware(http.HandlerFunc(UpdateOrder))).Methods("PUT")
	router.Router.HandleFunc("/order/delete/", AuthMiddleware(http.HandlerFunc(DelOrder))).Methods("DELETE")
	return router.Router
}

func NewServer() *Server {
	return &Server{}
}
