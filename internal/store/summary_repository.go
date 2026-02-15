package store

import (
	"database/sql"

	"pulse/internal/models"
)

// SummaryRepository handles database operations for AI-generated summaries.
type SummaryRepository struct {
	db *sql.DB
}

func NewSummaryRepository(db *sql.DB) *SummaryRepository {
	return &SummaryRepository{db: db}
}

// Insert stores a batch summary in the database.
func (r *SummaryRepository) Insert(lookbackHours, articleCount int, summary, sentiment string) error {
	_, err := r.db.Exec(`
		INSERT INTO summaries (lookback_hours, article_count, summary, sentiment)
		VALUES (?, ?, ?, ?)
	`, lookbackHours, articleCount, summary, sentiment)
	return err
}

// FetchLatest returns the most recent summary.
func (r *SummaryRepository) FetchLatest() (*models.Summary, error) {
	row := r.db.QueryRow(`
		SELECT id, lookback_hours, article_count, summary, sentiment, created_at
		FROM summaries
		ORDER BY created_at DESC
		LIMIT 1
	`)

	var s models.Summary
	err := row.Scan(&s.ID, &s.LookbackHours, &s.ArticleCount, &s.Summary, &s.Sentiment, &s.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
