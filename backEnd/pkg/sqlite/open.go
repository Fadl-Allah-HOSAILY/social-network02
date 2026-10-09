package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "pkg/database/social-network.db?_foreign_keys=on&_journal_mode=WAL")

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
