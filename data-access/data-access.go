package main

import (
	"database/sql"
	"fmt"
	"golangkah/data-access/bookwriter"
	"log"

	_ "github.com/lib/pq"
)

type Book struct {
	Id          string
	Title       string
	Pages       int64
	Writer      string
	Genre       string
	Description string
}

func main() {
	var connStr string = "postgres://maulanazn:t00r123@localhost/books?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")

	albums, err := bookwriter.BookByWriter("Tere Liye")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Book found: %v\n", albums)
}
