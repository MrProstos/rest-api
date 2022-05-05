package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	userData := map[string]interface{}{"username": "vlad", "token": ""}

	strJSON, err := json.Marshal(userData)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.Post("http://localhost:9999/newuser", "application/json", bytes.NewBuffer(strJSON))
	if err != nil {
		log.Fatal(err)
	}

	defer req.Body.Close()
	b, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(b))

}
