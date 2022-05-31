package book

import (
	"context"
)

type BookEntity struct {
	repo Repository
}

func NewBookEntity(repo Repository) *BookEntity {
	return &BookEntity{
		repo: repo,
	}
}

func (ent *BookEntity) GetBookById(ctx context.Context, id int) (*BookDto, error) {
	book, err := ent.repo.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (ent *BookEntity) GetAllBooks(ctx context.Context) ([]*BookDto, error) {
	categories, err := ent.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (ent *BookEntity) CreateNewBook(ctx context.Context, book *BookDto) error {
	err := ent.repo.Save(ctx, book)

	if err != nil {
		return err
	}

	return nil
}

func (ent *BookEntity) UpdateBookById(ctx context.Context, id int, book *BookDto) error {
	err := ent.repo.Update(ctx, id, book)

	if err != nil {
		return err
	}

	return nil
}

func (ent *BookEntity) DeleteBookById(ctx context.Context, id int) error {
	err := ent.repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
