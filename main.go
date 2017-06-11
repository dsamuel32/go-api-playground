package main

import (
	"log"
	"net/http"
)

//https://thenewstack.io/make-a-restful-json-api-go/
func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
