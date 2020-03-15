package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Info\n")
	if err := json.NewEncoder(w).Encode(info()); err != nil {
		panic(err)
	}
}

func Led(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Led\n")
}
