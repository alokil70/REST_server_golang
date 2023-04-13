package book

import "context"

type Repository interface {
	Create(ctx context.Context, user *Book) error
	FindAll(ctx context.Context) (a []Book, err error)
	FindOne(ctx context.Context, id string) (Book, error)
	Update(ctx context.Context, user Book) error
	Delete(ctx context.Context, id string) error
}
