package db_test

import (
	"testing"

	"github.com/MrProstos/rest-api/db"
)

var (
	oper   = db.Operator{Username: "vlad1", Token: "token1"}
	client = db.Client{
		Id:        1,
		Id_string: "123632",
		Firstname: "vlad",
		Lastname:  "mikgin",
		Operator:  "MTC",
	}
	ord = db.Order{
		Client_id:    "123",
		Phone_number: 89531222389,
		Title:        "asd",
		To:           "Vlad",
		Body:         "Text",
		Status:       "send",
	}
)

func TestAdd(t *testing.T) {

	err := oper.Add(db.DB)
	if err != nil {
		t.Error(err)
	}

	if err := client.Add(db.DB); err != nil {
		t.Error(err)
	}

	if err := ord.Add(db.DB); err != nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {

}

func TestDel(t *testing.T) {

}
