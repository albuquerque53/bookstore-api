package repo

import (
	"bookstoreapi/internal/domain/author"
	"bookstoreapi/internal/infra/database"
	"context"
)

type AuthorRepo struct {
	db *database.DatabaseManager
}

func NewAuthorRepo() *AuthorRepo {
	return &AuthorRepo{
		db: database.NewDatabaseManager(),
	}
}

func (repo *AuthorRepo) Get(ctx context.Context, id int) (*author.AuthorDto, error) {
	row, err := repo.db.QueryRow(ctx, "select id, name from authors where id=?", id)

	if err != nil {
		return nil, err
	}

	dto := &author.AuthorDto{}
	err = row.Scan(&dto.ID, &dto.Name)

	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (repo *AuthorRepo) GetAll(ctx context.Context) ([]*author.AuthorDto, error) {
	rows, err := repo.db.Query(ctx, "select id, name from authors")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	dtos := make([]*author.AuthorDto, 0)

	for rows.Next() {
		dto := &author.AuthorDto{}

		err = rows.Scan(&dto.ID, &dto.Name)

		if err != nil {
			return nil, err
		}

		dtos = append(dtos, dto)
	}

	return dtos, nil
}

func (repo *AuthorRepo) Save(ctx context.Context, dto *author.AuthorDto) error {
	_, err := repo.db.Query(ctx, "insert into authors(name) values(?)", dto.Name)

	return err
}
