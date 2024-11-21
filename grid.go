package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Grid struct {
	Matrix [][]string      `json:"matrix"`
	Words  map[string]bool `json:"words"`
}

type WordRequest struct {
	Word string `json:"word"`
}

func (cfg APIConfig) getGrid(w http.ResponseWriter, r *http.Request) {
	g := Grid{}
	createMatrix(&g, 10, 10)
	g.Words = map[string]bool{}
	respondWithJSON(w, http.StatusAccepted, g.Matrix)
}

func createMatrix(g *Grid, x int, y int) {
	g.Matrix = make([][]string, y)
	for i := 0; i < x; i++ {
		g.Matrix[i] = make([]string, y)
		for j := 0; j < y; j++ {
			g.Matrix[i][j] = "Z"
		}
	}
}

func (cfg APIConfig) handlerAddWord(w http.ResponseWriter, r *http.Request) {
	var body WordRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		log.Printf("error decoding json: %s\n", err)
		respondWithErorr(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	if body.Word == "" {
		log.Println("empty payload submitted")
		respondWithErorr(w, http.StatusBadRequest, "blank")
	}

	if _, ok := cfg.grid.Words[body.Word]; !ok {
		log.Println("duplicate word submitted")
		respondWithErorr(w, http.StatusBadRequest, "duplicate")
	}
	cfg.grid.Words[body.Word] = true
	respondWithJSON(w, http.StatusAccepted, cfg.grid.Words)
}
