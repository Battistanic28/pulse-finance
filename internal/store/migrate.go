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
	`

	_, err := db.Exec(schema)
	return err
}
