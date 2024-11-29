package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strings"
	"time"

	"github.com/aramirez3/wordsearch/internal/database"
	"github.com/google/uuid"
)

type Grid struct {
	Matrix      [][]string      `json:"matrix"`
	Words       map[string]bool `json:"words"`
	longestWord int
}

type WordsListRequest struct {
	Words []string `json:"words"`
}

func (cfg APIConfig) getGrid(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET GRID HANDLER")
	g := Grid{}
	createMatrix(&g, 10, 10)
	respondWithJSON(w, http.StatusAccepted, g.Matrix)
}

func (cfg *APIConfig) createGrid(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CREATE GRID")
	if len(cfg.grid.Words) == 0 {
		fmt.Printf("invalid empty payload")
		respondWithErorr(w, http.StatusBadRequest, "bad request")
		return
	}
	size := cfg.grid.longestWord + 5
	if size < 12 {
		size = 12
	}
	fmt.Printf("longest word: %d, grid size: %d x %d\n", cfg.grid.longestWord, size, size)
	createMatrix(&cfg.grid, size, size)
	for key := range cfg.grid.Words {
		cfg.addWordToMatrix(key)
	}

	data, err := json.Marshal(cfg.grid.Matrix)
	if err != nil {
		respondWithErorr(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	params := database.CreateGridParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Grid:      string(data),
	}
	grid, err := cfg.dbQueries.CreateGrid(r.Context(), params)
	if err != nil {
		respondWithErorr(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	respondWithJSON(w, http.StatusAccepted, grid)
}

func createMatrix(g *Grid, x int, y int) {
	g.Matrix = make([][]string, y)
	for i := 0; i < x; i++ {
		g.Matrix[i] = make([]string, y)
		for j := 0; j < y; j++ {
			g.Matrix[i][j] = "-"
		}
	}
}

func (cfg *APIConfig) addWordToMatrix(word string) {
	fmt.Printf("ADD WORD TO MATRIX: %s\n", word)
	cfg.insertWordInRow(word)
	// pick random direction and random forward or backward:
	//     left to right
	//     top to bottom
	//     diag up
	//     diag down
}

func (cfg *APIConfig) insertWordInRow(word string) {
	min := len(cfg.grid.Matrix[0]) - len(word)
	rowNum := rand.IntN(len(cfg.grid.Matrix[0])-min) + min
	cfg.grid.Matrix[rowNum] = leftToRight(cfg.grid.Matrix[rowNum], word)
}

func leftToRight(row []string, word string) []string {
	start := rand.IntN(len(row))
	if start < 0 {
		start = 0
	} else if start+len(word) > len(row) {
		start = len(row) - len(word) - 1
	}
	for i, char := range strings.Split(word, "") {
		row[start+i] = char
	}
	return row
}
