package storage

import (
	"fmt"
	// "os"
	// "path/filepath"
	"database/sql"
	// _ "modernc.org/sqlite"
)

type DB struct {
	conn *sql.DB
}

func Open(path string) (*DB, error) {
	fmt.Printf("Open database at path: %s\n", path)
	// os.MkdirAll(filepath.Dir(path), 0o755) // Creates the directory for the path if it doesn't exist
	return &DB{}, nil // Change this to return an actual database connection
}

func (db *DB) Close() error {
	fmt.Printf("Closing database connection\n")
	return nil // Change this to close the actual database connection
}

func (db *DB) initSchema() error {
	// Implement the database schema initialization logic here
	fmt.Printf("Initializing database schema\n")
	const schema = `
CREATE TABLE IF NOT EXISTS events (
	event_id TEXT PRIMARY KEY,
	guild_id TEXT NOT NULL,
	channel_id TEXT NOT NULL,
	name TEXT NOT NULL,
	start_time TEXT,
	end_time TEXT,
	status TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS event_subscriptions (
	event_id TEXT NOT NULL,
	user_id TEXT NOT NULL,
	subscribed_at TEXT NOT NULL,
	PRIMARY KEY (event_id, user_id)
);

CREATE TABLE IF NOT EXISTS voice_sessions (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id TEXT NOT NULL,
	user_id TEXT NOT NULL,
	channel_id TEXT NOT NULL,
	joined_at TEXT NOT NULL,
	left_at TEXT
);
`
	return nil // Change this to execute the actual schema initialization
}
