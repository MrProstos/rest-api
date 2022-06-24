package server

import (
	"encoding/json"
	"github.com/MrProstos/rest-api/internal/gateway/db"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func parseDate(str string) (string, error) {
	date, err := time.Parse("2006-01-02", str)
	if err != nil {
		return "", err
	}
	str = date.Format("2006-01-02")
	return str, nil
}

func ShowClients(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	client := db.NewClient()
	client.PhoneNum = vars["phone_num"]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}

func AddClient(w http.ResponseWriter, r *http.Request) {

}

func UpdateClient(w http.ResponseWriter, r *http.Request) {

}

func DelClient(w http.ResponseWriter, r *http.Request) {

}

func ShowOrder(w http.ResponseWriter, r *http.Request) {

}

func AddOrder(w http.ResponseWriter, r *http.Request) {

}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {

}

func DelOrder(w http.ResponseWriter, r *http.Request) {

}
