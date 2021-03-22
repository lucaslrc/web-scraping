package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	googlesearch "github.com/rocketlaunchr/google-search"
)

type JSONRequest struct {
	Search string `json:"search"`
}

type JSONResponse struct {
	Search string `json:"search"`
}

type Result struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func scrap(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}

	var jsonRequestBody JSONRequest

	err = json.Unmarshal(body, &jsonRequestBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}

	fmt.Println(googlesearch.Search(nil, "dolar"))

	w.WriteHeader(200)
}
