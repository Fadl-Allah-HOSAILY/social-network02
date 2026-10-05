package main

import (
	"backend/pkg/db/sqlite"
	"backend/route"
	"log"
	"net/http"
)

func main() {
	db, err := sqlite.OpenDb()
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlite.Migrate(db); err != nil {
		log.Fatal(err)
	}

	mux := route.NewRouter()

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
