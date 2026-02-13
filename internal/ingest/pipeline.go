package ingest

import (
	"encoding/json"

	"pulse/internal/models"
	"pulse/internal/store"
)

// Process decodes API response and writes to the database.
func Process(raw []byte, repo *store.NewsRepository) error {
	var resp models.APIResponse

	if err := json.Unmarshal(raw, &resp); err != nil {
		return err
	}

	for _, item := range resp.Feed {
		err := repo.Insert(store.News{
			Title:         item.Title,
			URL:           item.URL,
			Summary:       item.Summary,
			TimePublished: item.TimePublished,
			Sentiment:     item.OverallSentimentScore,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
