package user

import (
	"context"
	"strings"
	"test_RESTserver_01/internal/author"
	"test_RESTserver_01/pkg/client/postgresql"
	"test_RESTserver_01/pkg/logging"

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


// Create implements author.Repository
func (*repository) Create(ctx context.Context, user *author.Author) error {
	panic("unimplemented")
}

// Delete implements author.Repository
func (*repository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindAll implements author.Repository
func (*repository) FindAll(ctx context.Context) (a []author.Author, err error) {
	panic("unimplemented")
}

// FindOne implements author.Repository
func (*repository) FindOne(ctx context.Context, id string) (author.Author, error) {
	panic("unimplemented")
}

// Update implements author.Repository
func (*repository) Update(ctx context.Context, user author.Author) error {
	panic("unimplemented")
}
