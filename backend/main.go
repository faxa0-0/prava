package main

import (
	"log"
	"net/http"

	"github.com/faxa0-0/prava/config"
	"github.com/faxa0-0/prava/handlers"
	"github.com/faxa0-0/prava/services"
	"github.com/faxa0-0/prava/storage/postgres"
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
