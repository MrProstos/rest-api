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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := client.Add(db.DB); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	} else {
		fmt.Fprint(w, "Succses!")
	}
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	client := new(db.Client)
	err := decoder.Decode(&client)
	if err != nil {
		log.Println(err)

	}

}

func EnterOrder(w http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, err.Error())
	}
	var order db.Order

	err = json.Unmarshal(msg, &order)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, err.Error())
	}
	if err := order.Add(db.DB); err != nil {
		log.Println(err)
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, "Succses!")
	}
}
