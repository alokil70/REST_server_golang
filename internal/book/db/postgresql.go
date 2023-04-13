package book

import (
	"context"
	"fmt"
	"strings"
	"test_RESTserver_01/internal/author"
	"test_RESTserver_01/internal/book"
	"test_RESTserver_01/pkg/client/postgresql"
	"test_RESTserver_01/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}


func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) FindAll(ctx context.Context) (a []book.Book, err error) {
	q := `
		SELECT id, name FROM public.book
	`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	books := make([]book.Book, 0)

	for rows.Next() {
		var b book.Book

		err = rows.Scan(&b.ID, &b.Name)
		if err != nil {
			return nil, err
		}

		sq := `
			SELECT a.id, a.name
			FROM book_authors
			JOIN public.author a on a.id = book_authors.author_id
			WHERE book_id = $1;
		`

		authorRows, err := r.client.Query(ctx, sq)
		if err != nil {
			return nil, err
		}

		authors := make([]author.Author, 0)
		for authorRows.Next() {
			var a author.Author
	
			err = authorRows.Scan(&a.ID, &a.Name)
			if err != nil {
				return nil, err
			}
			
			authors = append(authors, a)
		}
		b.Author = authors

		books = append(books, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
