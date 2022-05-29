package category

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, id int) (*CategoryDto, error)
	GetAll(ctx context.Context) ([]*CategoryDto, error)
	Save(ctx context.Context, dto *CategoryDto) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, dto *CategoryDto) error
}
