package usecase

import (
	"errors"
	"testing"

	"github.com/evandersondev/test-golang-todo-list/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewFindAllUseCase(t *testing.T) {
	repository := &TodoRepositoryMock{}
	sut := NewFindAllUseCase(repository)
	assert.IsType(t, &FindAllUseCase{}, sut, "Expected sut to be of type *FindAllUseCase")
}

func TestFindAllUseCase_Execute_Success(t *testing.T) {
	repository := &TodoRepositoryMock{}
	repository.On("FindAll").Return([]entity.Todo{}, nil).Once()
	sut := NewFindAllUseCase(repository)

	todos, err := sut.Execute()

	assert.NoError(t, err)
	assert.NotNil(t, todos)
	repository.AssertExpectations(t)
}

func TestFindAllUseCase_Execute_Failure(t *testing.T) {
	repository := &TodoRepositoryMock{}
	repository.On("FindAll").Return([]entity.Todo{}, errors.New("failed to find todos")).Once()
	sut := NewFindAllUseCase(repository)

	todos, err := sut.Execute()
	assert.Error(t, err)
	assert.Empty(t, todos)
	repository.AssertExpectations(t)
}
