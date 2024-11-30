package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/aramirez3/wordsearch/internal/database"
)

type APIServer struct {
	addr      string
	apiConfig APIConfig
}

type APIConfig struct {
	dbQueries *database.Queries
	grid      Grid
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: ":" + addr,
		apiConfig: APIConfig{
			dbQueries: database.New(db),
			grid: Grid{
				Matrix:      [][]string{},
				Words:       map[string]bool{},
				longestWord: 0,
			},
		},
	}
}

func (s *APIServer) Start() error {
	router := http.NewServeMux()

	router.Handle("/", http.FileServer(http.Dir("static")))
	router.HandleFunc("/new", handlerNewWordForm)

	v1 := http.NewServeMux()
	v1.HandleFunc("POST /words", ValidateRequestMiddleware(s.apiConfig.handlerAddWord))
	v1.HandleFunc("DELETE /words", ValidateRequestMiddleware(s.apiConfig.handlerRemoveWord))
	v1.HandleFunc("GET /grids/{id}", s.apiConfig.getGrid)
	v1.HandleFunc("POST /grids", s.apiConfig.createGrid)
	v1.HandleFunc("GET /grids", s.apiConfig.getGrids)

	router.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	middlewareChain := MiddlewareChain(
		RequestLoggerMiddleware,
		// RequireAuthMiddleware,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
	}

	fmt.Printf("addr: %v\n", server.Addr)

	log.Printf("Server started on %s", s.addr)

	return server.ListenAndServe()
}
