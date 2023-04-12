package author

import (
	"context"
	"fmt"
	"strings"
	"test_RESTserver_01/internal/author"
	"test_RESTserver_01/pkg/client/postgresql"
	"test_RESTserver_01/pkg/logging"

	"github.com/jackc/pgx/v5/pgconn"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) author.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) Create(ctx context.Context, author *author.Author) error {
	q := `
		INSERT INTO
		author (name)
		VALUES ($1)
		RETURNING id
	`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	if err := r.client.QueryRow(ctx, q, author.Name).Scan(&author.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			e := fmt.Errorf(fmt.Sprintf("SQL error: %s, Detail: %s, Where: %s, Code: %s, SqlState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Logger.Error(e)
			return e
		}
		return err
	}

	return nil
}

func (r *repository) FindAll(ctx context.Context) (a []author.Author, err error) {
	q := `
		SELECT id, name FROM public.author
	`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	authors := make([]author.Author, 0)

	for rows.Next() {
		var a author.Author

		err = rows.Scan(&a.ID, &a.Name)
		if err != nil {
			return nil, err
		}

		authors = append(authors, a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (author.Author, error) {
	q := `
		SELECT id, name FROM public.author WHERE id = $1
	`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var a author.Author
	err := r.client.QueryRow(ctx, q, id).Scan(&a.ID, &a.Name)
	if err != nil {
		return author.Author{}, err
	}
	
	return a, nil
}

func (r *repository) Update(ctx context.Context, user author.Author) error {
	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	return nil	
}
