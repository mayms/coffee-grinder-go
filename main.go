package main

import (
	"log"
	"net/http"
)

func main() {
	initGpio()
	r := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
