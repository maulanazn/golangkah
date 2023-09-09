package bookwriter

import (
	"database/sql"
	"fmt"
	"golangkah/model"
	"log"

	_ "github.com/lib/pq"
)

// type Book struct {
// 	Id          string
// 	Title       string
// 	Pages       int64
// 	Writer      string
// 	Genre       string
// 	Description string
// }

func BookByWriter(writer string) ([]model.Book, error) {
	var book []model.Book

	var connStr string = "postgres://maulanazn:t00r123@localhost/books?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM book WHERE writer = $1", writer)

	if err != nil {
		return nil, fmt.Errorf("bookByWriter %q: %v", writer, err)
	}

	defer rows.Close()

	for rows.Next() {
		var bk model.Book
		if err := rows.Scan(&bk.Id, &bk.Title, &bk.Pages, &bk.Writer, &bk.Genre, &bk.Description); err != nil {
			return nil, fmt.Errorf("bookyByWriter %q: %v", writer, err)
		}

		book = append(book, bk)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("bookByArtist %q: %v", writer, err)
	}

	return book, nil
}
