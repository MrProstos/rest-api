package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAddClient(t *testing.T) {
	userData := map[string]interface{}{
		"phone_num": "111",
		"firstname": "test",
		"lastname":  "test",
		"birthday":  "1-1-2000",
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
	if string(b) != "Succses!" {
		t.Error(string(b))
	} else {
		fmt.Println(string(b))
	}

}

func TestUpdateClient(t *testing.T) {
	userData := map[string]interface{}{
		"id":        1,
		"phone_num": "777",
		"firstname": "vlad",
		"lastname":  "mikhin",
		"birthday":  "22.02.22",
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
	if string(b) != "Succses!" {
		t.Error(string(b))
	} else {
		fmt.Println(string(b))
	}
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
	if string(b) != "Succses!" {
		t.Error(string(b))
	} else {
		fmt.Println(string(b))
	}
}
