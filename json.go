package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error marshaling json: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	_, err = w.Write(responseData)
	if err != nil {
		log.Printf("error sending json: %s", err)
	}
}
