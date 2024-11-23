package main

import (
	"fmt"
	"net/http"
)

type Grid struct {
	Matrix [][]string      `json:"matrix"`
	Words  map[string]bool `json:"words"`
}

func (cfg APIConfig) getGrid(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET GRID HANDLER")
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
