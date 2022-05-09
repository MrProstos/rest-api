package db_test

import (
	"fmt"
	"testing"

	"github.com/MrProstos/rest-api/db"
)

func TestAdd(t *testing.T) {
	oper := db.Operator{Username: "vlad1", Token: "token1"}

	err := oper.Add(db.DB)
	if err != nil {
		t.Error(err)
	}

	client := db.Client{
		Id:        1,
		Id_string: "123632",
		Firstname: "vlad",
		Lastname:  "mikgin",
		Operator:  "MTC",
	}
	if err := client.Add(db.DB); err != nil {
		t.Error(err)
	}

	ord := db.Order{
		Client_id: "123",
		Title:     "asd",
		To:        "Vlad",
		Body:      "Text",
		Status:    "send",
	}
	if err := ord.Add(db.DB); err != nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {
	oper := db.Operator{
		Username: "vlad1",
	}
	client := db.Client{
		Id_string: "123632",
	}
	ord := db.Order{
		Client_id: "123",
		Title:     "asd",
	}
	fmt.Println(oper, client, ord)

}

func TestDel(t *testing.T) {

}
