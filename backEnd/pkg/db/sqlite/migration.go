package sqlite

import (
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

func Migrate(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "pkg/db/migrations/sqlite",
	}

	mg, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		return err
	}

	fmt.Println("migrations applied:", mg)
	return nil
}
