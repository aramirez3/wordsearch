package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respondError(w http.ResponseWriter, code int, e string) {
	type errorResponse struct {
		Error string
	}
	respondWithJSON(w, code, errorResponse{
		Error: e,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("marshaling payload: %v\n", payload)
	pay, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error marshaling json: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	_, err = w.Write(pay)
	if err != nil {
		fmt.Printf("error sending json: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
