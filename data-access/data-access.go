package main

import (
	"database/sql"
	"fmt"
	"golangkah/data-access/bookbyid"
	"golangkah/data-access/bookwriter"
	"log"

	_ "github.com/lib/pq"
)

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

	album, err := bookbyid.BookById("6fa1d6e6-52c6-4a79-9aff-1b03ceb73f55")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Book found: %v\n", album)
}
