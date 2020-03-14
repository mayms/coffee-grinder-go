package main

import (
	"fmt"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Info\n")
}

func Led(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Led\n")
}
