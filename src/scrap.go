package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	googlesearch "github.com/rocketlaunchr/google-search"
)

type JSONRequest struct {
	Search string `json:"search"`
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

	itens := []googlesearch.Result{}

	reqResult, err := googlesearch.Search(nil, jsonRequestBody.Search)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}

	for i := 0; i < len(reqResult); i++ {
		if len(reqResult[i].Title) != 0 {
			itens = append(itens, reqResult[i])
		}
	}

	resp, err := json.Marshal(itens)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}

	resp, err = json.MarshalIndent(itens, "", " ")

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}

	w.Write(resp)
}
