package repo

import (
	"bookstoreapi/internal/domain/dto"
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

func (repo *AuthorRepo) Get(ctx context.Context, id int) (*dto.AuthorDto, error) {
	row, err := repo.db.QueryRow(ctx, "select id, name from users where id=?", id)

	if err != nil {
		return nil, err
	}

	dto := &dto.AuthorDto{}
	err = row.Scan(&dto.ID, &dto.Name)

	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (repo *AuthorRepo) GetAll(ctx context.Context) ([]*dto.AuthorDto, error) {
	rows, err := repo.db.Query(ctx, "select id, name from users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	dtos := make([]*dto.AuthorDto, 0)

	for rows.Next() {
		dto := &dto.AuthorDto{}

		err = rows.Scan(&dto.ID, &dto.Name)

		if err != nil {
			return nil, err
		}

		dtos = append(dtos, dto)
	}

	return dtos, nil
}
