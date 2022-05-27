package server

import (
	"encoding/json"
	"fmt"
	"github.com/MrProstos/rest-api/internal/gateway/db"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ShowClients(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	client := new(db.Client)
	client.Phone_num = vars["phone_num"]

	err := db.ManageDb.Show(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Error(w, fmt.Sprint(client), http.StatusOK)
}

func AddClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	client := new(db.Client)
	err := decoder.Decode(&client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.ManageDb.Add(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Error(w, "Success", http.StatusOK)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	client := new(db.Client)
	if err := decoder.Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := db.ManageDb.Update(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Error(w, "Success", http.StatusOK)
}

func DelClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	client := new(db.Client)
	err := decoder.Decode(&client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.ManageDb.Del(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Error(w, "Success", http.StatusOK)
}

func ShowOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ord := new(db.Order)
	strClientId := vars["client_id"]

	intClientId, err := strconv.Atoi(strClientId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ord.Client_id = uint(intClientId)

	err = db.ManageDb.Show(ord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Error(w, fmt.Sprint(ord), http.StatusOK)

}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	order := new(db.Order)
	err := decoder.Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.ManageDb.Add(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Error(w, "Success", http.StatusOK)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	order := new(db.Order)
	err := decoder.Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.ManageDb.Update(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Error(w, "Success", http.StatusOK)

}

func DelOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	order := new(db.Order)
	err := decoder.Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.ManageDb.Del(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Error(w, "Success", http.StatusOK)
}
