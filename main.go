package main

import (
	"fmt"

	"github.com/MrProstos/rest-api/db"
)

func main() {
	o := db.Operator{Username: "MTC", Token: "123"}
	s := db.Client{
		Id:        "1",
		Firstname: "Test",
		Lastname:  "asd",
		Operator:  o.Username,
	}
	fmt.Println(s)
}
