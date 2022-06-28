package main

import (
	"github.com/MrProstos/rest-api/internal/server"
	"log"
	"net/http"
)

func main() {

	err := http.ListenAndServe(":2000", server.NewServer().Init())
	if err != nil {
		log.Fatal(err)
	}
}
