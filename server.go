package main

import (
	"log"
	"net/http"

	"github.com/ishuah/batian/api"
)

func main() {
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":2906", router))
}
