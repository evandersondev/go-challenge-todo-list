package usecase

import (
	"testing"

	"github.com/evandersondev/test-golang-todo-list/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewFindTodoById(t *testing.T) {
	repository := &TodoRepositoryMock{}
	sut := NewFindTodoById(repository)
	assert.IsType(t, &FindTodoById{}, sut, "Expected sut to be of type *FindTodoById")
}

func TestNewFindTodoById_Execute_Success(t *testing.T) {
	repository := &TodoRepositoryMock{}
	expectedTodo := &entity.Todo{
		ID:          "1",
		Title:       "Test Title",
		Description: "Test Description",
		Done:        false,
	}

	repository.On("FindById", "1").Return(expectedTodo, nil)
	sut := NewFindTodoById(repository)

	todo, err := sut.Execute("1")

	assert.NoError(t, err)
	assert.Equal(t, expectedTodo, todo)
	repository.AssertExpectations(t)
}

func TestNewFindTodoById_Execute_Error(t *testing.T) {
	repository := &TodoRepositoryMock{}
	repository.On("FindById", "invalid-id").Return(nil, assert.AnError)
	sut := NewFindTodoById(repository)

	todo, err := sut.Execute("invalid-id")

	assert.Error(t, err)
	assert.Nil(t, todo)
	repository.AssertExpectations(t)
}
