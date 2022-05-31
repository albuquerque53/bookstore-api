package book

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBookById(t *testing.T) {
	type sceneries struct {
		ctx        context.Context
		id         int
		expectBook *BookDto
		expectErr  error
	}

	testSceneries := []sceneries{
		{ctx: context.Background(), id: 1, expectBook: buildExpectDto(1), expectErr: nil},
		{ctx: context.Background(), id: 2, expectBook: nil, expectErr: errors.New("test error")},
	}

	for _, scenery := range testSceneries {
		mock := new(RepoMock)

		mock.On("Get", scenery.ctx, scenery.id).Return(scenery.expectBook, scenery.expectErr)

		ent := getBookEntity(mock)

		dto, err := ent.GetBookById(scenery.ctx, scenery.id)

		assert.Equal(t, scenery.expectErr, err)
		assert.Equal(t, scenery.expectBook, dto)
	}
}

func TestGetAllBooks(t *testing.T) {
	type sceneries struct {
		ctx         context.Context
		expectBooks []*BookDto
		expectErr   error
	}

	dtos := []*BookDto{
		buildExpectDto(1),
		buildExpectDto(2),
	}

	testSceneries := []sceneries{
		{ctx: context.Background(), expectBooks: dtos, expectErr: nil},
		{ctx: context.Background(), expectBooks: nil, expectErr: errors.New("test error")},
	}

	for _, scenery := range testSceneries {
		mock := new(RepoMock)

		mock.On("GetAll", scenery.ctx).Return(scenery.expectBooks, scenery.expectErr)

		ent := getBookEntity(mock)

		dto, err := ent.GetAllBooks(scenery.ctx)

		assert.Equal(t, scenery.expectErr, err)
		assert.Equal(t, scenery.expectBooks, dto)
	}
}

func TestCreateNewBook(t *testing.T) {
	type sceneries struct {
		ctx       context.Context
		category  BookDto
		expectErr error
	}

	testSceneries := []sceneries{
		{ctx: context.Background(), category: *buildExpectDto(1), expectErr: nil},
		{ctx: context.Background(), category: *buildExpectDto(2), expectErr: errors.New("test error")},
	}

	for _, scenery := range testSceneries {
		mock := new(RepoMock)

		mock.On("Save", scenery.ctx, &scenery.category).Return(scenery.expectErr)

		ent := getBookEntity(mock)

		err := ent.CreateNewBook(scenery.ctx, &scenery.category)

		assert.Equal(t, scenery.expectErr, err)
	}
}

func TestDeleteBookById(t *testing.T) {
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

		ent := getBookEntity(mock)

		err := ent.DeleteBookById(scenery.ctx, scenery.id)

		assert.Equal(t, scenery.expectErr, err)
	}
}

func TestUpdateBookById(t *testing.T) {
	type sceneries struct {
		ctx       context.Context
		id        int
		category  *BookDto
		expectErr error
	}

	testSceneries := []sceneries{
		{ctx: context.Background(), id: 1, category: buildExpectDto(1), expectErr: nil},
		{ctx: context.Background(), id: 2, category: buildExpectDto(2), expectErr: errors.New("test error")},
	}

	for _, scenery := range testSceneries {
		mock := new(RepoMock)

		mock.On("Update", scenery.ctx, scenery.id, scenery.category).Return(scenery.expectErr)

		ent := getBookEntity(mock)

		err := ent.UpdateBookById(scenery.ctx, scenery.id, scenery.category)

		assert.Equal(t, scenery.expectErr, err)
	}
}

func getBookEntity(mock *RepoMock) *BookEntity {
	return &BookEntity{
		repo: mock,
	}
}

func buildExpectDto(id int) *BookDto {
	return &BookDto{ID: id, Title: "test_book", AuthorID: 1, CategoryID: 2, CreatedAt: "2022-05-29 11:46:50", UpdatedAt: "2022-05-29 11:46:50"}
}
