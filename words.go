package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type WordsList struct {
	Words map[string]bool `json:"words"`
}

type WordRequest struct {
	Word string `json:"word"`
}

func handlerNewWordForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NEW WORD FORM")
	template := `<div>
	<input type="text" name="word" </>
	<button
		hx-post="api/v1/words"
		hx-trigger="click"
		hx-ext="json-enc"
		>Add word</button>
</div>
`
	w.Write([]byte(template))
}

func (cfg *APIConfig) handlerAddWord(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ADD WORD HANDLER")
	payload := WordRequest{}
	err := cfg.validateWordRequest(r, &payload)
	if err != nil {
		respondWithErorr(w, http.StatusBadRequest, err.Error())
		return
	}

	if _, ok := cfg.grid.Words[payload.Word]; ok {
		respondWithErorr(w, http.StatusBadRequest, "duplicate")
		return
	}

	if len(payload.Word) > cfg.grid.longestWord {
		cfg.grid.longestWord = len(payload.Word)
	}
	cfg.grid.Words[payload.Word] = true
	respondWithJSON(w, http.StatusAccepted, cfg.grid.Words)
}

func (cfg APIConfig) handlerRemoveWord(w http.ResponseWriter, r *http.Request) {
	fmt.Println("REMOVE WORD HANDLER")
	payload := WordRequest{}
	err := cfg.validateWordRequest(r, &payload)
	if err != nil {
		respondWithErorr(w, http.StatusBadRequest, err.Error())
		return
	}

	delete(cfg.grid.Words, payload.Word)
	respondWithJSON(w, http.StatusAccepted, cfg.grid.Words)
}

func (cfg *APIConfig) validateWordRequest(r *http.Request, payload *WordRequest) error {
	fmt.Println("VALIDATE WORDS PAYLOAD")
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&payload); err != nil {
		log.Printf("error decoding json: %s\n", err)
		return fmt.Errorf("bad request")
	}

	if payload.Word == "" {
		log.Println("invalid empty payload")
		return fmt.Errorf("bad request")
	}

	alphaRegEx := regexp.MustCompile("^[a-zA-Z]+$")
	if !alphaRegEx.MatchString(payload.Word) {
		log.Println("only alpha chars are allowed")
		return fmt.Errorf("bad request")
	}

	return nil
}
