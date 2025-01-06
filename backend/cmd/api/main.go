package main

import (
	"log"
	"net/http"

	"github.com/faxa0-0/prava/internal/config"
	"github.com/faxa0-0/prava/internal/handlers"
	"github.com/faxa0-0/prava/internal/services"
	"github.com/faxa0-0/prava/internal/storage/postgres"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	storage, err := postgres.NewPostgresQuizStorage(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}
	srv := services.NewQuizService(storage)
	h := handlers.NewQuizHandler(srv)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /quiz", h.GetQuizHandler)

	log.Printf("Starting api on localhost:%s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, mux))
}
