package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type WordsResponse struct {
	Words []string `json:"words"`
}

type Words struct {
	List map[string]bool `json:"words_map"`
}

type matrix = [][]string

type newGrid struct {
	W     int
	H     int
	Words []string
}

type wordPayload struct {
	Word string `json:"word"`
}

var words = Words{
	List: map[string]bool{},
}

var wordsResponse = WordsResponse{}

func handlerNewGrid(w http.ResponseWriter, r *http.Request) {
	grid := newGrid{}
	_, err := json.Marshal(&grid)
	if err != nil {
		respondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	g := generateGrid(grid)
	respondWithJSON(w, http.StatusCreated, g)
}

func generateGrid(g newGrid) matrix {
	fmt.Printf("Create new grid with these attribues: %v\n", g)
	m := make(matrix, g.W)
	for i := range m {
		m[i] = make([]string, g.H)
	}
	return m
}

func handlerNew(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/base.html", "templates/new.html")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func addWord(w http.ResponseWriter, r *http.Request) {
	returnPage(w, "new")
}

func addWordApi(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request body: %v\n", r.Body)
	err := r.ParseForm()
	if err != nil {
		fmt.Println("error parsing form")
		return
	}
	fmt.Printf("parsed form data: %v\n", r.PostForm)
	postWord, ok := r.PostForm["word"]
	if !ok || postWord[0] == "" {
		return
	}
	word := postWord[0]
	_, ok = words.List[word]
	if !ok {
		body := wordPayload{
			Word: word,
		}
		words.List[word] = true
		wordsResponse.Words = append(wordsResponse.Words, word)
		fmt.Printf("wordsResponse: %v\n", body)
	}
	respondWithJSON(w, http.StatusAccepted, wordsResponse)
}
