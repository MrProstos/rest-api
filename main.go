package main

import (
	"net/http"

	"github.com/MrProstos/rest-api/server"
	"github.com/MrProstos/rest-api/utils"
	"github.com/gorilla/mux"
)

func main() {

	utils.InitLogger()
	utils.Logger.Info("Start server!")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/addclient/", server.AddClient).Methods("POST")
	router.HandleFunc("/updateclient/", server.UpdateClient).Methods("PUT")
	router.HandleFunc("/delclient/", server.DelClient).Methods("DELETE")
	router.HandleFunc("/addorder/", server.AddOrder).Methods("POST")
	router.HandleFunc("/updateorder/", server.UpdateOrder).Methods("PUT")
	router.HandleFunc("/delorder/", server.DelOrder).Methods("DELETE")
	err := http.ListenAndServe(":2000", router)
	if err != nil {
		utils.Logger.Error(err.Error())
	}

}
