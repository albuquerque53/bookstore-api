package author

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthorById(t *testing.T) {
	type sceneries struct {
		ctx          context.Context
		id           int
		expectAuthor *AuthorDto
		expectErr    error
	}

	testSceneries := []sceneries{
		{ctx: context.Background(), id: 1, expectAuthor: buildExpectDto(1), expectErr: nil},
		{ctx: context.Background(), id: 2, expectAuthor: nil, expectErr: errors.New("test error")},
	}

	for _, scenery := range testSceneries {
		mock := new(RepoMock)

		mock.On("Get", scenery.ctx, scenery.id).Return(scenery.expectAuthor, scenery.expectErr)

		ent := getAuthorEntity(mock)

		dto, err := ent.GetAuthorById(scenery.ctx, scenery.id)

		assert.Equal(t, scenery.expectErr, err)
		assert.Equal(t, scenery.expectAuthor, dto)
	}
}

func TestGetAllAuthors(t *testing.T) {
	type sceneries struct {
		ctx          context.Context
		expectAuthor []*AuthorDto
		expectErr    error
	}

	dtos := []*AuthorDto{
		buildExpectDto(1),
		buildExpectDto(2),
	}

	testSceneries := []sceneries{
		{ctx: context.Background(), expectAuthor: dtos, expectErr: nil},
		{ctx: context.Background(), expectAuthor: nil, expectErr: errors.New("test error")},
	}

	for _, scenery := range testSceneries {
		mock := new(RepoMock)

		mock.On("GetAll", scenery.ctx).Return(scenery.expectAuthor, scenery.expectErr)

		ent := getAuthorEntity(mock)

		dto, err := ent.GetAllAuthors(scenery.ctx)

		assert.Equal(t, scenery.expectErr, err)
		assert.Equal(t, scenery.expectAuthor, dto)
	}
}

func TestCreateNewAuthor(t *testing.T) {
	type sceneries struct {
		ctx       context.Context
		author    AuthorDto
		expectErr error
	}

	testSceneries := []sceneries{
		{ctx: context.Background(), author: *buildExpectDto(1), expectErr: nil},
		{ctx: context.Background(), author: *buildExpectDto(2), expectErr: errors.New("test error")},
	}

	for _, scenery := range testSceneries {
		mock := new(RepoMock)

		mock.On("Save", scenery.ctx, &scenery.author).Return(scenery.expectErr)

		ent := getAuthorEntity(mock)

		err := ent.CreateNewAuthor(scenery.ctx, &scenery.author)

		assert.Equal(t, scenery.expectErr, err)
	}
}

func TestDeleteAuthorById(t *testing.T) {
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

		ent := getAuthorEntity(mock)

		err := ent.DeleteAuthorById(scenery.ctx, scenery.id)

		assert.Equal(t, scenery.expectErr, err)
	}
}

func TestUpdateAuthorById(t *testing.T) {
	type sceneries struct {
		ctx       context.Context
		id        int
		author    *AuthorDto
		expectErr error
	}

	testSceneries := []sceneries{
		{ctx: context.Background(), id: 1, author: buildExpectDto(1), expectErr: nil},
		{ctx: context.Background(), id: 2, author: buildExpectDto(2), expectErr: errors.New("test error")},
	}

	for _, scenery := range testSceneries {
		mock := new(RepoMock)

		mock.On("Update", scenery.ctx, scenery.id, scenery.author).Return(scenery.expectErr)

		ent := getAuthorEntity(mock)

		err := ent.UpdateAuthorById(scenery.ctx, scenery.id, scenery.author)

		assert.Equal(t, scenery.expectErr, err)
	}
}

func getAuthorEntity(mock *RepoMock) *AuthorEntity {
	return &AuthorEntity{
		repo: mock,
	}
}

func buildExpectDto(id int) *AuthorDto {
	return &AuthorDto{ID: id, Name: "test_author", CreatedAt: "2022-05-29 11:46:50", UpdatedAt: "2022-05-29 11:46:50"}
}
