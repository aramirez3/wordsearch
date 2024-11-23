package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type WordRequest struct {
	Word string `json:"word"`
}

func handlerNewWordForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NEW WORD FORM")
	template := `<div>
	<input type="text" name="word" </>
	<button
		hx-post="words"
		hx-trigger="click"
		>Add word</button>
</div>
`
	w.Write([]byte(template))
}

func (cfg APIConfig) handlerAddWord(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ADD WORD HANDLER")
	word, err := cfg.validateWordRequest(w, r.Body)
	if err != nil {
		respondWithErorr(w, http.StatusBadRequest, err.Error())
	}
	log.Printf("Saving word: %s\n", word)
	cfg.grid.Words[word] = true
	respondWithJSON(w, http.StatusAccepted, cfg.grid.Words)
}

func (cfg *APIConfig) validateWordRequest(w http.ResponseWriter, payload io.ReadCloser) (string, error) {
	var body WordRequest
	decoder := json.NewDecoder(payload)
	if err := decoder.Decode(&body); err != nil {
		log.Printf("error decoding json: %s\n", err)
		respondWithErorr(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	if body.Word == "" {
		log.Println("empty payload submitted")
		return "", fmt.Errorf("blank")
	}

	if _, ok := cfg.grid.Words[body.Word]; !ok {
		log.Println("duplicate word submitted")
		return "", fmt.Errorf("duplicate")
	}
	return body.Word, nil
}
