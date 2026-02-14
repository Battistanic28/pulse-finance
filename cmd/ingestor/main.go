package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"pulse/internal/client"
	"pulse/internal/ingest"
	"pulse/internal/store"
)

func main() {
	_ = godotenv.Load()

	if len(os.Args) < 3 {
		log.Fatal("Usage: go run ./cmd/ingestor <TICKER> <LOOKBACK_HOURS>")
	}

	ticker := os.Args[1]
	lookbackStr := os.Args[2]

	lookbackHours, err := strconv.Atoi(lookbackStr)
	if err != nil {
		log.Fatal("Invalid lookback hours")
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

	repo := store.NewNewsRepository(db)

	apiKey := os.Getenv("ALPHAVANTAGE_API_KEY")
	if apiKey == "" {
		log.Fatal("ALPHAVANTAGE_API_KEY environment variable is required")
	}

	av := client.Client{
		APIKey:     apiKey,
		HTTPClient: nil,
	}

	raw, err := av.Fetch(ticker, lookbackHours)
	if err != nil {
		log.Fatal(err)
	}

	if err := ingest.Process(raw, repo); err != nil {
		log.Fatal(err)
	}

	log.Println("Ingestion complete")
}
