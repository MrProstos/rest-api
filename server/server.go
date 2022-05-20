package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MrProstos/rest-api/db"
)

func AddClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	client := new(db.Client)
	err := decoder.Decode(&client)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Db_manage.Add(client); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	client := new(db.Client)
	if err := decoder.Decode(&client); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Db_manage.Update(client); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)

}

func DelClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	client := new(db.Client)
	if err := decoder.Decode(&client); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Db_manage.Del(client); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	order := new(db.Order)
	err := decoder.Decode(&order)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Db_manage.Add(order); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	order := new(db.Order)
	if err := decoder.Decode(&order); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Db_manage.Update(order); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)

}

func DelOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	order := new(db.Order)
	if err := decoder.Decode(&order); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Db_manage.Del(order); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}
