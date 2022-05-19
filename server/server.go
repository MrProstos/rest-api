package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := client.Update(); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, http.StatusOK)

}

func EnterOrder(w http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, err.Error())
	}

	var order db.Order

	if err := json.Unmarshal(msg, &order); err != nil {
		log.Println(err)
		fmt.Fprint(w, err.Error())
	}

	if err := order.Add(); err != nil {
		log.Println(err)
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, "Succses!")
	}
}
