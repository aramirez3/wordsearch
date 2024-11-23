package main

import (
	"fmt"
	"log"
	"net/http"
)

type APIServer struct {
	addr      string
	apiConfig APIConfig
}

type APIConfig struct {
	grid Grid
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr:      ":" + addr,
		apiConfig: APIConfig{},
	}
}

func (s *APIServer) Start() error {
	router := http.NewServeMux()

	router.Handle("/", http.FileServer(http.Dir("static")))

	v1 := http.NewServeMux()
	v1.HandleFunc("GET /grids/{id}", s.apiConfig.getGrid)
	v1.HandleFunc("POST /words", s.apiConfig.handlerAddWord)

	router.Handle("/api/v1/", http.StripPrefix("/api/v1", RequireAuthMiddleware(v1)))

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
