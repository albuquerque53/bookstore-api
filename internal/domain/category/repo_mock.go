package category

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

func (m *RepoMock) Get(ctx context.Context, id int) (*CategoryDto, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(*CategoryDto), args.Error(1)
}

func (m *RepoMock) GetAll(ctx context.Context) ([]*CategoryDto, error) {
	args := m.Called(ctx)

	return args.Get(0).([]*CategoryDto), args.Error(1)
}

func (m *RepoMock) Save(ctx context.Context, dto *CategoryDto) error {
	args := m.Called(ctx, dto)

	return args.Error(0)
}

func (m *RepoMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)

	return args.Error(0)
}

func (m *RepoMock) Update(ctx context.Context, id int, dto *CategoryDto) error {
	args := m.Called(ctx, id, dto)

	return args.Error(0)
}
