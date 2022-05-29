package author

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

func (m *RepoMock) Get(ctx context.Context, id int) (*AuthorDto, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(*AuthorDto), args.Error(1)
}

func (m *RepoMock) GetAll(ctx context.Context) ([]*AuthorDto, error) {
	args := m.Called(ctx)

	return args.Get(0).([]*AuthorDto), args.Error(1)
}

func (m *RepoMock) Save(ctx context.Context, dto *AuthorDto) error {
	args := m.Called(ctx, dto)

	return args.Error(0)
}

func (m *RepoMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)

	return args.Error(0)
}

func (m *RepoMock) Update(ctx context.Context, id int, dto *AuthorDto) error {
	args := m.Called(ctx, id, dto)

	return args.Error(0)
}
