package book

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, id int) (*BookDto, error)
	GetAll(ctx context.Context) ([]*BookDto, error)
	Save(ctx context.Context, dto *BookDto) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, dto *BookDto) error
}
