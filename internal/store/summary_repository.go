package store

import "database/sql"

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
