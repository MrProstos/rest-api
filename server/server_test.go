package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/MrProstos/rest-api/db"
)

func TestAddClient(t *testing.T) {
	userData := map[string]interface{}{
		"Phone_num": "252",
		"Firstname": "test",
		"Lastname":  "test",
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
		"Client_id": 4,
		"Phone_num": "3234",
		"Firstname": "vlad",
		"Lastname":  "mikhin",
	}

	strJSON, err := json.Marshal(userData)
	if err != nil {
		t.Error(err)
	}

	olddata, newdata := new(db.Client), new(db.Client)
	fmt.Println(newdata)
	err = json.Unmarshal(strJSON, &olddata)
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
