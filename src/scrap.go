package main

import (
	"log"
	"net/http"
)

func scrap(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("chegou")
	w.WriteHeader(200)
}
