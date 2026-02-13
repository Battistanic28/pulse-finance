package store

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// Open creates or opens the SQLite database.
func Open(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Performance & concurrency settings
	pragmas := []string{
		"PRAGMA journal_mode = WAL;",   // allow concurrent reads
		"PRAGMA synchronous = NORMAL;", // faster writes
		"PRAGMA foreign_keys = ON;",
	}

	for _, p := range pragmas {
		if _, err := db.Exec(p); err != nil {
			return nil, err
		}
	}

	return db, nil
}
