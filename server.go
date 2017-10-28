package main

import (
	"log"
	"net/http"

	"github.com/ishuah/batian/handlers"
)

func main() {
	router := handlers.NewRouter()
	log.Fatal(http.ListenAndServe(":2906", router))
}
