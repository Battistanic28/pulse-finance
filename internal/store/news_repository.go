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

// FetchSince returns all news articles published after the given time string.
// The timeFrom parameter should match the time_published format (e.g. "20060102T1504").
func (r *NewsRepository) FetchSince(timeFrom string) ([]News, error) {
	rows, err := r.db.Query(`
		SELECT title, url, summary, time_published, sentiment_score
		FROM news
		WHERE time_published >= ?
		ORDER BY time_published DESC
	`, timeFrom)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []News
	for rows.Next() {
		var n News
		if err := rows.Scan(&n.Title, &n.URL, &n.Summary, &n.TimePublished, &n.Sentiment); err != nil {
			return nil, err
		}
		articles = append(articles, n)
	}
	return articles, rows.Err()
}
