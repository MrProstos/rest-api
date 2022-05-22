package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

type TestClient struct {
	ID        uint
	Phone_num string
	Firstname string
	Lastname  string
	Birthday  string
	//OrderID   uint
	Orders []TestOrder
}

type TestOrder struct {
	Client_id uint
	Title     string
	To        string
	Body      string
	Status    uint
}

func TestShowClietn(t *testing.T) {
	body, err := http.Get("http://localhost:2000/showclietns/777")
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(body.Body)
	fmt.Println(string(b))
}

func TestAddClient(t *testing.T) {
	o := TestOrder{
		Title:  "Vlad",
		To:     "Vlad",
		Body:   "Test",
		Status: 0,
	}
	array := []TestClient{{
		Phone_num: "777",
		Firstname: "Vlad",
		Lastname:  "Mikhin",
		Birthday:  "1999-07-22",
		//OrderID:   1,
		Orders: []TestOrder{o},
	}, {
		Phone_num: "666",
		Firstname: "Test",
		Lastname:  "Test",
		Birthday:  "00-00-0000",
	}, {
		Phone_num: "123456789",
		Firstname: "QQQQ",
		Lastname:  "ZZZZ",
		Birthday:  "0000000",
	}}

	for i := range array {

		strJSON, err := json.Marshal(array[i])
		if err != nil {
			t.Error(err)
		}

		body, err := http.Post("http://localhost:2000/addclient/", "application/json", bytes.NewBuffer(strJSON))
		if err != nil {
			t.Error(err)
		}
		b, _ := ioutil.ReadAll(body.Body)

		if string(b) != "200" {
			t.Error(string(b))
		}
		fmt.Printf("Тест номер %v %v\n", i+1, string(b))
	}
}

func TestUpdateClient(t *testing.T) {
	array := []TestClient{{
		ID:        1,
		Phone_num: "UPDAT1",
		Firstname: "UPDATE",
		Lastname:  "UPDATE",
		Birthday:  "UPDATE",
	}, {
		ID:        2,
		Phone_num: "UPDATE2",
		Firstname: "UPDATE",
		Lastname:  "UPDATE",
		Birthday:  "UPDATE",
	}, {
		ID:        3,
		Phone_num: "UPDATE3",
		Firstname: "UPDATE",
		Lastname:  "UPDATE",
		Birthday:  "UPDATE",
	}}

	for i := range array {

		strJSON, err := json.Marshal(array[i])
		if err != nil {
			t.Error(err)
		}

		body, err := http.Post("http://localhost:2000/updateclient/", "application/json", bytes.NewBuffer(strJSON))
		if err != nil {
			t.Error(err)
		}

		b, _ := ioutil.ReadAll(body.Body)

		if string(b) != "200" {
			t.Error(string(b))
		}
		fmt.Printf("Тест номер %v %v\n", i+1, string(b))
	}
}

func TestDelClietn(t *testing.T) {
	array := []TestClient{{
		ID: 1,
	}, {
		ID: 2,
	}, {
		ID: 3,
	}}

	for i := range array {

		strJSON, err := json.Marshal(array[i])
		if err != nil {
			t.Error(err)
		}

		body, err := http.Post("http://localhost:2000/delclient/", "application/json", bytes.NewBuffer(strJSON))
		if err != nil {
			t.Error(err)
		}

		b, _ := ioutil.ReadAll(body.Body)

		if string(b) != "200" {
			t.Error(string(b))
		}
		fmt.Printf("Тест номер %v %v\n", i+1, string(b))
	}

}

func TestAddOrder(t *testing.T) {
}

func TestUpdateOrder(t *testing.T) {
}
