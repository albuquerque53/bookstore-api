package repo

import (
	"bookstoreapi/internal/domain/book"
	"bookstoreapi/internal/infra/database"
	"context"
)

type BookRepo struct {
	db *database.DatabaseManager
}

func NewBookRepo() *BookRepo {
	return &BookRepo{
		db: database.NewDatabaseManager(),
	}
}

func (repo *BookRepo) Get(ctx context.Context, id int) (*book.BookDto, error) {
	row, err := repo.db.QueryRow(ctx, "select id, title, author_id, category_id, created_at, updated_at from books where id=?", id)

	if err != nil {
		return nil, err
	}

	dto := &book.BookDto{}
	err = row.Scan(&dto.ID, &dto.Title, &dto.AuthorID, &dto.CategoryID, &dto.CreatedAt, &dto.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (repo *BookRepo) GetAll(ctx context.Context) ([]*book.BookDto, error) {
	rows, err := repo.db.Query(ctx, "select id, title, author_id, category_id, created_at, updated_at from books")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	dtos := make([]*book.BookDto, 0)

	for rows.Next() {
		dto := &book.BookDto{}

		err = rows.Scan(&dto.ID, &dto.Title, &dto.AuthorID, &dto.CategoryID, &dto.CreatedAt, &dto.UpdatedAt)

		if err != nil {
			return nil, err
		}

		dtos = append(dtos, dto)
	}

	return dtos, nil
}

func (repo *BookRepo) Save(ctx context.Context, dto *book.BookDto) error {
	_, err := repo.db.Query(ctx, "insert into books(title, author_id, category_id) values(?, ?, ?)", dto.Title, dto.AuthorID, dto.CategoryID)

	return err
}

func (repo *BookRepo) Delete(ctx context.Context, id int) error {
	_, err := repo.db.Query(ctx, "delete from books where id = ?", id)

	return err
}

func (repo *BookRepo) Update(ctx context.Context, id int, dto *book.BookDto) error {
	bookDto, err := repo.Get(ctx, id)

	if err != nil {
		return err
	}

	if dto.Title != "" {
		bookDto.Title = dto.Title
	}

	if dto.AuthorID != 0 {
		bookDto.AuthorID = dto.AuthorID
	}

	if dto.CategoryID != 0 {
		bookDto.CategoryID = dto.CategoryID
	}

	_, err = repo.db.Query(ctx, "update books set title = ?, author_id = ?, category_id = ? where id = ?", bookDto.Title, bookDto.AuthorID, bookDto.CategoryID, id)

	return err
}
