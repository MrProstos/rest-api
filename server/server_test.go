package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestEnterClient(t *testing.T) {
	userData := map[string]interface{}{
		"Phone_num": "89208823212",
		"Firstname": "vlad",
		"Lastname":  "mikhin",
		"Birthday":  "22-07-1999",
	}
	strJSON, err := json.Marshal(userData)
	if err != nil {
		t.Error(err)
	}
	body, err := http.Post("http://localhost:2000/addclient/", "application/json", bytes.NewBuffer(strJSON))
	if err != nil {
		t.Error(err)
	}
	b, _ := ioutil.ReadAll(body.Body)
	fmt.Println(string(b))

}

func TestEnterOrder(t *testing.T) {
	orderData := map[string]interface{}{
		"Client_id": 3,
	}

	strJSON, err := json.Marshal(orderData)
	if err != nil {
		t.Error(err)
	}

	body, err := http.Post("http://localhost:2000/addorder/", "application/json", bytes.NewBuffer(strJSON))
	if err != nil {
		t.Error(err)
	}
	b, _ := ioutil.ReadAll(body.Body)
	fmt.Println(string(b))
}
