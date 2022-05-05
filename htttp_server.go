package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Userdata struct {
	Username string
	Token    string
}

func newuser_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		user := Userdata{}
		json.Unmarshal(body, &user)
		user.Token = uuid.New().String()
		fmt.Fprintf(w, "%v %v", user.Username, user.Token)
	}
}

func main() {
	log.Println("Starting HTTP server on port 9999")
	http.HandleFunc("/newuser", newuser_handler)
	http.ListenAndServe("localhost:9999", nil)
}
