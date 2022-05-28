package author

import (
	"context"
)

type AuthorEntity struct {
	repo Repository
}

func NewAuthorEntity(repo Repository) *AuthorEntity {
	return &AuthorEntity{
		repo: repo,
	}
}

func (ent *AuthorEntity) GetAuthorById(ctx context.Context, id int) (*AuthorDto, error) {
	author, err := ent.repo.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return author, nil
}

func (ent *AuthorEntity) GetAllAuthors(ctx context.Context) ([]*AuthorDto, error) {
	authors, err := ent.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (ent *AuthorEntity) CreateNewAuthor(ctx context.Context, author *AuthorDto) error {
	err := ent.repo.Save(ctx, author)

	if err != nil {
		return err
	}

	return nil
}

func (ent *AuthorEntity) DeleteAuthorById(ctx context.Context, id int) error {
	err := ent.repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
