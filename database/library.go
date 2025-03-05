package database

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/juhitha20/golang/module/module"
)

func CreateBook(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var book module.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		return err
	}
	query := "INSERT INTO books(title, author, year) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, book.Title, book.Author, book.Year)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(book)
}

func GetBooks(db *sql.DB) ([]module.Book, error) {
	rows, err := db.Query("SELECT title, author, year FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []module.Book
	for rows.Next() {
		var book module.Book
		if err := rows.Scan(&book.Title, &book.Author, &book.Year); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
