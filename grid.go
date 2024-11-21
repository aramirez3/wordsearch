package main

import (
	"net/http"
)

type Grid struct {
	Matrix [][]string `json:"matrix"`
	Words  []string   `json:"words"`
}

func (cfg APIConfig) getGrid(w http.ResponseWriter, r *http.Request) {
	g := Grid{}
	g.createMatrix(10, 10)
	g.Words = []string{"hello", "world"}
	respondWithJSON(w, http.StatusAccepted, g.Matrix)
}

func (g *Grid) createMatrix(x int, y int) {
	g.Matrix = make([][]string, y)
	for i := 0; i < x; i++ {
		g.Matrix[i] = make([]string, y)
		for j := 0; j < y; j++ {
			g.Matrix[i][j] = "Z"
		}
	}
}
