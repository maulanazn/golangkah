package bookbyid

import (
	"database/sql"
	"fmt"
	"golangkah/model"
	"log"

	_ "github.com/lib/pq"
)

func BookById(id string) (model.Book, error) {
	var connStr string = "postgres://maulanazn:t00r123@localhost/books?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Failed to connect")
	}

	var book model.Book

	row := db.QueryRow("SELECT * FROM book WHERE id = ?", id)
	if err := row.Scan(&book.Id, &book.Title, &book.Pages, &book.Writer, &book.Genre, &book.Description); err != nil {
		if err == sql.ErrNoRows {
			return book, fmt.Errorf("BookById %s: no such book", id)
		}
		return book, fmt.Errorf("BookById: %s: %v", id, err)
	}

	return book, nil
}
