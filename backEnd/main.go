package main

import (
	"backend/pkg/db/sqlite"
	"log"
)

func main() {
	db, err := sqlite.OpenDb()
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlite.Migrate(db); err != nil {
		log.Fatal(err)
	}
}