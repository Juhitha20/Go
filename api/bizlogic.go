package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/juhitha20/golang/module/database"
)

func CreateBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return database.CreateBook(db, w, r)
}

func GetBooksLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	books, err := database.GetBooks(db)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(books)
}
