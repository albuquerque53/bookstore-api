package category

import (
	"context"
)

type CategoryEntity struct {
	repo Repository
}

func NewCategoryEntity(repo Repository) *CategoryEntity {
	return &CategoryEntity{
		repo: repo,
	}
}

func (ent *CategoryEntity) GetCategoryById(ctx context.Context, id int) (*CategoryDto, error) {
	category, err := ent.repo.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (ent *CategoryEntity) GetAllCategories(ctx context.Context) ([]*CategoryDto, error) {
	categories, err := ent.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (ent *CategoryEntity) CreateNewCategory(ctx context.Context, category *CategoryDto) error {
	err := ent.repo.Save(ctx, category)

	if err != nil {
		return err
	}

	return nil
}

func (ent *CategoryEntity) UpdateCategoryById(ctx context.Context, id int, category *CategoryDto) error {
	err := ent.repo.Update(ctx, id, category)

	if err != nil {
		return err
	}

	return nil
}

func (ent *CategoryEntity) DeleteCategoryById(ctx context.Context, id int) error {
	err := ent.repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
