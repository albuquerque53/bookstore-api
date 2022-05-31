package category

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCategoryById(t *testing.T) {
	type sceneries struct {
		ctx            context.Context
		id             int
		expectCategory *CategoryDto
		expectErr      error
	}

	testSceneries := []sceneries{
		{ctx: context.Background(), id: 1, expectCategory: buildExpectDto(1), expectErr: nil},
		{ctx: context.Background(), id: 2, expectCategory: nil, expectErr: errors.New("test error")},
	}

	for _, scenery := range testSceneries {
		mock := new(RepoMock)

		mock.On("Get", scenery.ctx, scenery.id).Return(scenery.expectCategory, scenery.expectErr)

		ent := getCategoryEntity(mock)

		dto, err := ent.GetCategoryById(scenery.ctx, scenery.id)

		assert.Equal(t, scenery.expectErr, err)
		assert.Equal(t, scenery.expectCategory, dto)
	}
}

func TestGetAllCategories(t *testing.T) {
	type sceneries struct {
		ctx              context.Context
		expectCategories []*CategoryDto
		expectErr        error
	}

	dtos := []*CategoryDto{
		buildExpectDto(1),
		buildExpectDto(2),
	}

	testSceneries := []sceneries{
		{ctx: context.Background(), expectCategories: dtos, expectErr: nil},
		{ctx: context.Background(), expectCategories: nil, expectErr: errors.New("test error")},
	}

	for _, scenery := range testSceneries {
		mock := new(RepoMock)

		mock.On("GetAll", scenery.ctx).Return(scenery.expectCategories, scenery.expectErr)

		ent := getCategoryEntity(mock)

		dto, err := ent.GetAllCategories(scenery.ctx)

		assert.Equal(t, scenery.expectErr, err)
		assert.Equal(t, scenery.expectCategories, dto)
	}
}

func TestCreateNewCategory(t *testing.T) {
	type sceneries struct {
		ctx       context.Context
		category  CategoryDto
		expectErr error
	}

	testSceneries := []sceneries{
		{ctx: context.Background(), category: *buildExpectDto(1), expectErr: nil},
		{ctx: context.Background(), category: *buildExpectDto(2), expectErr: errors.New("test error")},
	}

	for _, scenery := range testSceneries {
		mock := new(RepoMock)

		mock.On("Save", scenery.ctx, &scenery.category).Return(scenery.expectErr)

		ent := getCategoryEntity(mock)

		err := ent.CreateNewCategory(scenery.ctx, &scenery.category)

		assert.Equal(t, scenery.expectErr, err)
	}
}

func TestDeleteCategoryById(t *testing.T) {
	type sceneries struct {
		ctx       context.Context
		id        int
		expectErr error
	}

	testSceneries := []sceneries{
		{ctx: context.Background(), id: 1, expectErr: nil},
		{ctx: context.Background(), id: 2, expectErr: errors.New("test error")},
	}

	for _, scenery := range testSceneries {
		mock := new(RepoMock)

		mock.On("Delete", scenery.ctx, scenery.id).Return(scenery.expectErr)

		ent := getCategoryEntity(mock)

		err := ent.DeleteCategoryById(scenery.ctx, scenery.id)

		assert.Equal(t, scenery.expectErr, err)
	}
}

func TestUpdateCategoryById(t *testing.T) {
	type sceneries struct {
		ctx       context.Context
		id        int
		category  *CategoryDto
		expectErr error
	}

	testSceneries := []sceneries{
		{ctx: context.Background(), id: 1, category: buildExpectDto(1), expectErr: nil},
		{ctx: context.Background(), id: 2, category: buildExpectDto(2), expectErr: errors.New("test error")},
	}

	for _, scenery := range testSceneries {
		mock := new(RepoMock)

		mock.On("Update", scenery.ctx, scenery.id, scenery.category).Return(scenery.expectErr)

		ent := getCategoryEntity(mock)

		err := ent.UpdateCategoryById(scenery.ctx, scenery.id, scenery.category)

		assert.Equal(t, scenery.expectErr, err)
	}
}

func getCategoryEntity(mock *RepoMock) *CategoryEntity {
	return &CategoryEntity{
		repo: mock,
	}
}

func buildExpectDto(id int) *CategoryDto {
	return &CategoryDto{ID: id, Name: "test_category", CreatedAt: "2022-05-29 11:46:50", UpdatedAt: "2022-05-29 11:46:50"}
}
