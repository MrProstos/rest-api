package server

import (
	"encoding/json"
	"github.com/MrProstos/rest-api/internal/gateway/db"
	"github.com/gorilla/mux"
	"net/http"
)

func ShowClients(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	client := db.NewClient()
	client.PhoneNum = vars["phone_num"]

	_, err := db.Tables.Select(client)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, &client)
}

func AddClient(w http.ResponseWriter, r *http.Request) {
	client := db.NewClient()

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&client)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}

	_, err = db.Tables.Insert(client)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, &client)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	client := db.NewClient()

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&client)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}

	_, err = db.Tables.Update(client)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, &client)
}

func DelClient(w http.ResponseWriter, r *http.Request) {
	client := db.NewClient()

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&client)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}

	_, err = db.Tables.Delete(client)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, &client)
}

func ShowOrder(w http.ResponseWriter, r *http.Request) {
	order := db.NewOrder()

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&order)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}

	_, err = db.Tables.Select(order)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, &order)
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	order := db.NewOrder()

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&order)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}

	_, err = db.Tables.Insert(order)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, &order)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	order := db.NewOrder()

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&order)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}

	_, err = db.Tables.Update(order)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, &order)
}

func DelOrder(w http.ResponseWriter, r *http.Request) {
	order := db.NewOrder()

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&order)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}

	_, err = db.Tables.Delete(order)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, &order)
}
