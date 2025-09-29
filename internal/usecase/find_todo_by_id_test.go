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
	repository.On("FindById", "1").Return(entity.Todo{}, nil)
	sut := NewFindTodoById(repository)

	_, _ = sut.repository.FindById("1")

	// assert.Error(t, err)
	// assert.NotNil(t, todo)
	// repository.AssertExpectations(t)
}
