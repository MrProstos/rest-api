package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MrProstos/rest-api/db"
	"github.com/MrProstos/rest-api/utils"
	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Login")
}

func ShowClients(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	client := new(db.Client)
	client.Phone_num = vars["phone_num"]

	if err := db.Db_manage.Show(client); err != nil {
		utils.Logger.Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%v %v", http.StatusOK, client)
}

func AddClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	client := new(db.Client)
	err := decoder.Decode(&client)
	if err != nil {
		utils.Logger.Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Db_manage.Add(client); err != nil {
		utils.Logger.Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	client := new(db.Client)
	if err := decoder.Decode(&client); err != nil {
		utils.Logger.Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Db_manage.Update(client); err != nil {
		utils.Logger.Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)

}

func DelClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	client := new(db.Client)
	if err := decoder.Decode(&client); err != nil {
		utils.Logger.Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Db_manage.Del(client); err != nil {
		utils.Logger.Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}

func ShowOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ord := new(db.Order)
	str_client_id := vars["client_id"]

	int_client_id, err := strconv.Atoi(str_client_id)
	if err != nil {
		utils.Logger.Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ord.Client_id = uint(int_client_id)

	if err := db.Db_manage.Show(ord); err != nil {
		utils.Logger.Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%v %v", http.StatusOK, ord)

}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	order := new(db.Order)
	err := decoder.Decode(&order)
	if err != nil {
		utils.Logger.Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Db_manage.Add(order); err != nil {
		utils.Logger.Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	order := new(db.Order)
	if err := decoder.Decode(&order); err != nil {
		utils.Logger.Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Db_manage.Update(order); err != nil {
		utils.Logger.Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)

}

func DelOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	order := new(db.Order)
	if err := decoder.Decode(&order); err != nil {
		utils.Logger.Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Db_manage.Del(order); err != nil {
		utils.Logger.Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}
