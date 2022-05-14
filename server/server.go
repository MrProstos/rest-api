package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MrProstos/rest-api/db"
)

func EnterClient(w http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, err.Error())
	}
	var client db.Client

	err = json.Unmarshal(msg, &client)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, err.Error())
	}
	if err := client.Add(db.DB); err != nil {
		log.Println(err)
		fmt.Fprint(w, err.Error())
	}
}
