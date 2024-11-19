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
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr:      ":" + addr,
		apiConfig: APIConfig{},
	}
}

func (s *APIServer) Start() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /api/v1/grids/{id}", s.apiConfig.getGrid)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	middlewareChain := MiddlewareChain(
		RequestLoggerMiddleware,
		RequireAuthMiddleware,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
	}

	fmt.Printf("addr: %v\n", server.Addr)

	log.Printf("Server started on %s", s.addr)

	return server.ListenAndServe()
}
