package models

// Summary represents a batch AI-generated summary of news articles.
type Summary struct {
	ID            int64
	LookbackHours int
	ArticleCount  int
	Summary       string
	Sentiment     string
	CreatedAt     string
}
