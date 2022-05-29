package repo

import (
	"bookstoreapi/internal/domain/category"
	"bookstoreapi/internal/infra/database"
	"context"
)

type CategoryRepo struct {
	db *database.DatabaseManager
}

func NewCategoryRepo() *CategoryRepo {
	return &CategoryRepo{
		db: database.NewDatabaseManager(),
	}
}

func (repo *CategoryRepo) Get(ctx context.Context, id int) (*category.CategoryDto, error) {
	row, err := repo.db.QueryRow(ctx, "select id, name, created_at, updated_at from categories where id=?", id)

	if err != nil {
		return nil, err
	}

	dto := &category.CategoryDto{}
	err = row.Scan(&dto.ID, &dto.Name, &dto.CreatedAt, &dto.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (repo *CategoryRepo) GetAll(ctx context.Context) ([]*category.CategoryDto, error) {
	rows, err := repo.db.Query(ctx, "select id, name, created_at, updated_at from categories")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	dtos := make([]*category.CategoryDto, 0)

	for rows.Next() {
		dto := &category.CategoryDto{}

		err = rows.Scan(&dto.ID, &dto.Name, &dto.CreatedAt, &dto.UpdatedAt)

		if err != nil {
			return nil, err
		}

		dtos = append(dtos, dto)
	}

	return dtos, nil
}

func (repo *CategoryRepo) Save(ctx context.Context, dto *category.CategoryDto) error {
	_, err := repo.db.Query(ctx, "insert into categories(name) values(?)", dto.Name)

	return err
}

func (repo *CategoryRepo) Delete(ctx context.Context, id int) error {
	_, err := repo.db.Query(ctx, "delete from categories where id = ?", id)

	return err
}

func (repo *CategoryRepo) Update(ctx context.Context, id int, dto *category.CategoryDto) error {
	_, err := repo.db.Query(ctx, "update categories set name = ? where id = ?", dto.Name, id)

	return err
}
