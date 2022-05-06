package db_test

import (
	"fmt"
	"testing"

	"github.com/MrProstos/rest-api/db"
)

func TestNewOperator(t *testing.T) {
	s, err := db.NewOperator("test", "token")
	fmt.Println(s, err)
}

func TestOperator(t *testing.T) {
	str, err := db.NewOperator("vlad1", "token1")
	if err != nil {
		t.Error(err)
	}

	err = str.Add(db.DB)
	if err != nil {
		t.Error(err)
	}

	err = str.Del(db.DB) // Проверить на счет нескольких удалений
	if err != nil {
		t.Error(err)
	}
}
