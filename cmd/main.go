package main

import (
	"fmt"
	"github.com/MrProstos/rest-api/internal/server"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server starting!")
	err := http.ListenAndServe(":2000", server.GetRouter())
	if err != nil {
		log.Fatal(err)
	}
}
