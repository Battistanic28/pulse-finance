package analyze

import (
	"fmt"
	"log"
	"strings"
	"time"

	"pulse/internal/client"
	"pulse/internal/store"
)

// Run fetches recent articles from the DB, sends them to Ollama for
// summarization, and stores the resulting summary.
func Run(newsRepo *store.NewsRepository, summaryRepo *store.SummaryRepository, ollama *client.OllamaClient, lookbackHours int) error {
	timeFrom := time.Now().UTC().Add(-time.Duration(lookbackHours) * time.Hour).Format("20060102T1504")

	articles, err := newsRepo.FetchSince(timeFrom)
	if err != nil {
		return fmt.Errorf("fetch articles: %w", err)
	}

	if len(articles) == 0 {
		log.Println("No articles found in the given time range")
		return nil
	}

	log.Printf("Found %d articles, sending to Ollama for analysis...", len(articles))

	prompt := buildPrompt(articles)

	response, err := ollama.Generate(prompt)
	if err != nil {
		return fmt.Errorf("ollama generate: %w", err)
	}

	// Extract sentiment line from response (last line starting with "SENTIMENT:")
	summary, sentiment := parseSummaryResponse(response)

	if err := summaryRepo.Insert(lookbackHours, len(articles), summary, sentiment); err != nil {
		return fmt.Errorf("store summary: %w", err)
	}

	log.Println("Analysis complete")
	fmt.Println("\n--- Summary ---")
	fmt.Println(summary)
	fmt.Printf("\nSentiment: %s\n", sentiment)
	fmt.Printf("Articles analyzed: %d\n", len(articles))

	return nil
}

func buildPrompt(articles []store.News) string {
	var b strings.Builder

	b.WriteString("You are a financial news analyst. Below are recent news articles with their sentiment scores.\n")
	b.WriteString("Provide a concise summary of the key themes and market-moving events.\n")
	b.WriteString("At the end, on a new line, write exactly: SENTIMENT: <bullish|bearish|neutral|mixed>\n\n")

	for i, a := range articles {
		fmt.Fprintf(&b, "Article %d:\n", i+1)
		fmt.Fprintf(&b, "Title: %s\n", a.Title)
		fmt.Fprintf(&b, "Summary: %s\n", a.Summary)
		fmt.Fprintf(&b, "Sentiment Score: %.4f\n\n", a.Sentiment)
	}

	b.WriteString("Provide your analysis now.")
	return b.String()
}

func parseSummaryResponse(response string) (summary, sentiment string) {
	lines := strings.Split(strings.TrimSpace(response), "\n")

	for i := len(lines) - 1; i >= 0; i-- {
		if strings.HasPrefix(strings.TrimSpace(lines[i]), "SENTIMENT:") {
			sentiment = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(lines[i]), "SENTIMENT:"))
			summary = strings.TrimSpace(strings.Join(lines[:i], "\n"))
			return summary, sentiment
		}
	}

	// If no SENTIMENT line found, return everything as summary
	return strings.TrimSpace(response), "unknown"
}
