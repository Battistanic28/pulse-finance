package store

import "database/sql"

// Migrate creates database tables if they do not exist.
func Migrate(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS news (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		url TEXT UNIQUE,
		summary TEXT,
		time_published TEXT,
		sentiment_score REAL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_news_time
	ON news(time_published);

	CREATE TABLE IF NOT EXISTS summaries (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		lookback_hours INTEGER,
		article_count INTEGER,
		summary TEXT,
		sentiment TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(schema)
	return err
}
