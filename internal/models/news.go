package models

// APIResponse matches Alpha Vantage NEWS_SENTIMENT response.
type APIResponse struct {
	Feed []FeedItem `json:"feed"`
}

type FeedItem struct {
	Title                 string  `json:"title"`
	URL                   string  `json:"url"`
	Summary               string  `json:"summary"`
	TimePublished         string  `json:"time_published"`
	OverallSentimentScore float64 `json:"overall_sentiment_score"`
}
