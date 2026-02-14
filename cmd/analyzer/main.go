package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"pulse/internal/analyze"
	"pulse/internal/client"
	"pulse/internal/store"
)

func main() {
	_ = godotenv.Load()

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run ./cmd/analyzer <LOOKBACK_HOURS>")
	}

	lookbackHours, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Invalid lookback hours")
	}

	ollamaURL := os.Getenv("OLLAMA_URL")
	if ollamaURL == "" {
		ollamaURL = "http://localhost:11434"
	}

	ollamaModel := os.Getenv("OLLAMA_MODEL")
	if ollamaModel == "" {
		ollamaModel = "llama3.1:8b"
	}

	dbPath := "data/news.db"
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		log.Fatal(err)
	}

	db, err := store.Open(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := store.Migrate(db); err != nil {
		log.Fatal(err)
	}

	newsRepo := store.NewNewsRepository(db)
	summaryRepo := store.NewSummaryRepository(db)

	ollama := &client.OllamaClient{
		BaseURL: ollamaURL,
		Model:   ollamaModel,
	}

	if err := analyze.Run(newsRepo, summaryRepo, ollama, lookbackHours); err != nil {
		log.Fatal(err)
	}
}
