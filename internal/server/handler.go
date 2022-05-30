package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MrProstos/rest-api/internal/gateway/db"
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

	client := new(db.Client)
	client.PhoneNum = vars["phone_num"]

	err := db.ManageDb.Show(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}

func AddClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	client := new(db.Client)
	err := decoder.Decode(&client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	str, err := parseDate(client.Birthday)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	client.Birthday = str

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

	if client.Birthday != "" {
		str, err := parseDate(client.Birthday)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		client.Birthday = str
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
	phoneNum := vars["phone_num"]

	ord.PhoneNum = phoneNum

	err := db.ManageDb.Show(ord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ord)
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
