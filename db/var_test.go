package db_test

import (
	"log"
	"testing"

	"github.com/MrProstos/rest-api/db"
)

func TestOperator(t *testing.T) {
	str := db.Operator{Username: "vlad1", Token: "token1"}

	err := str.Add(db.DB)
	if err != nil {
		t.Error(err)
	}

	err = str.Del(db.DB) // Проверить на счет нескольких удалений
	if err != nil {
		t.Error(err)
	}
}

func TestClient(t *testing.T) {
	str := db.Client{
		Id_string: "123632",
		Firstname: "vlad",
		Lastname:  "mikgin",
		Operator:  "MTC",
	}
	if err := str.Add(db.DB); err != nil {
		log.Fatal(err)
	}

}
