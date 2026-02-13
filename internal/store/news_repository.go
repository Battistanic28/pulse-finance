package store

import "database/sql"

// News represents a record stored in the database.
type News struct {
	Title         string
	URL           string
	Summary       string
	TimePublished string
	Sentiment     float64
}

// NewsRepository handles database operations for news.
type NewsRepository struct {
	db *sql.DB
}

func NewNewsRepository(db *sql.DB) *NewsRepository {
	return &NewsRepository{db: db}
}

// Insert adds a news article if it does not already exist.
func (r *NewsRepository) Insert(n News) error {
	_, err := r.db.Exec(`
		INSERT OR IGNORE INTO news
		(title, url, summary, time_published, sentiment_score)
		VALUES (?, ?, ?, ?, ?)
	`, n.Title, n.URL, n.Summary, n.TimePublished, n.Sentiment)

	return err
}
