package usecase

import (
	"github.com/evandersondev/test-golang-todo-list/internal/dto"
	"github.com/evandersondev/test-golang-todo-list/internal/entity"
	"github.com/stretchr/testify/mock"
)

type TodoRepositoryMock struct {
	mock.Mock
}

func (m *TodoRepositoryMock) Create(todo entity.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *TodoRepositoryMock) FindAll() ([]entity.Todo, error) {
	args := m.Called()
	return args.Get(0).([]entity.Todo), args.Error(1)
}

func (m *TodoRepositoryMock) FindById(id string) (*entity.Todo, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Todo), args.Error(1)
}

func (m *TodoRepositoryMock) Update(id string, dto dto.UpdateTodoDTO) (*entity.Todo, error) {
	args := m.Called(id, dto)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Todo), args.Error(1)
}

func (m *TodoRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
