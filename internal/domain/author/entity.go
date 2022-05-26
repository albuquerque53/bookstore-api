package author

import (
	"context"
	"log"
)

type AuthorEntity struct {
	repo Repository
}

func NewAuthorEntity(repo Repository) *AuthorEntity {
	return &AuthorEntity{
		repo: repo,
	}
}

func (ent *AuthorEntity) GetAuthorById(ctx context.Context, id int) *AuthorDto {
	author, err := ent.repo.Get(ctx, id)

	if err != nil {
		log.Fatal("error on author search")
	}

	return author
}

func (ent *AuthorEntity) GetAllAuthors(ctx context.Context) []*AuthorDto {
	authors, err := ent.repo.GetAll(ctx)

	if err != nil {
		log.Fatalf("error on authors list: %s", err)
	}

	return authors
}

func (ent *AuthorEntity) CreateNewAuthor(ctx context.Context, author *AuthorDto) {
	err := ent.repo.Save(ctx, author)

	if err != nil {
		log.Fatal("error on author create")
	}
}
