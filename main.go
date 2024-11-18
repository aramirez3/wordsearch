package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

var staticFiles embed.FS

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf(".env unreadable: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := staticFiles.Open("./static/index.html")
		if err != nil {
			fmt.Println("could not load f")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("found f, will copy now")
		defer f.Close()
		if _, err := io.Copy(w, f); err != nil {
			fmt.Println("could not copy f")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Post("/words", addWordApi)
	v1Router.Post("/grid", handlerNewGrid)
	v1Router.Post("/words/", addWord)

	v1Router.Get("/healthz", handlerReadiness)

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Addr:        ":" + port,
		Handler:     router,
		ReadTimeout: time.Duration(5 * time.Second),
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
