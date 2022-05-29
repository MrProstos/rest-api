package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MrProstos/rest-api/internal/server"
)

func main() {
	fmt.Println("Server starting!")
	err := http.ListenAndServe(":2000", server.GetRouter())
	if err != nil {
		log.Fatal(err)
	}
}
