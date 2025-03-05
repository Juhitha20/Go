package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/juhitha20/golang/module/api"
)

// test statement for pr
func main() {
	connStr := "user=juhitha dbname=mygolangdb password=Juhitha123# sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)

	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}
	fmt.Println("Connection established successfully!")

	api.RegisterRoutes(db)
	fmt.Println("Sent the request to routes")

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
