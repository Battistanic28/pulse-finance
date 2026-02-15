package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"pulse/internal/client"
	"pulse/internal/store"
)

func main() {
	_ = godotenv.Load()

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

	summaryRepo := store.NewSummaryRepository(db)

	summary, err := summaryRepo.FetchLatest()
	if err != nil {
		log.Fatalf("Failed to fetch latest summary: %v", err)
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailTo := os.Getenv("EMAIL_TO")

	if smtpHost == "" || smtpPort == "" || smtpUser == "" || smtpPassword == "" || emailFrom == "" || emailTo == "" {
		log.Fatal("Missing required SMTP environment variables (SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASSWORD, EMAIL_FROM, EMAIL_TO)")
	}

	subject := fmt.Sprintf("Pulse Finance Summary - %s", summary.CreatedAt)
	body := fmt.Sprintf("Market Summary\n\n%s\n\nSentiment: %s\nArticles analyzed: %d\nLookback: %d hours\nGenerated: %s\n",
		summary.Summary, summary.Sentiment, summary.ArticleCount, summary.LookbackHours, summary.CreatedAt)

	mailer := client.NewSMTPClient(smtpHost, smtpPort, emailFrom, smtpPassword, emailTo)

	if err := mailer.Send(subject, body); err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	log.Printf("Email sent to %s", emailTo)
}
