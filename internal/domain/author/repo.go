package author

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, id int) (*AuthorDto, error)
	GetAll(ctx context.Context) ([]*AuthorDto, error)
	Save(ctx context.Context, dto *AuthorDto) error
	Delete(ctx context.Context, id int) error
}
