package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type JSONRequest struct {
	Search string `json:"search"`
}

type JSONResponse struct {
	Search string `json:"search"`
}

func scrap(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
	var jsonBody JSONRequest
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
	// log.Println(jsonBody.Search)
	log.Println(string([]byte(body)))
	w.WriteHeader(200)
}
