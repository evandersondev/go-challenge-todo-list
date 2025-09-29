package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeleteTodoUseCase(t *testing.T) {
	repository := &TodoRepositoryMock{}
	sut := NewDeleteTodoUseCase(repository)
	assert.IsTypef(t, &DeleteTodoUseCase{}, sut, "Expected sut to be of type *DeleteTodoUseCase")
}

func TestDeleteTodoUseCase_Execute_Success(t *testing.T) {
	repository := &TodoRepositoryMock{}
	repository.On("Delete", "1").Return(nil).Once()
	sut := NewDeleteTodoUseCase(repository)

	err := sut.Execute("1")

	assert.NoError(t, err)
	repository.AssertExpectations(t)
}

func TestDeleteTodoUseCase_Execute_Failure(t *testing.T) {
	repository := &TodoRepositoryMock{}
	repository.On("Delete", "1").Return(errors.New("failed to delete todo")).Once()
	sut := NewDeleteTodoUseCase(repository)

	err := sut.Execute("1")

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to delete todo")
	repository.AssertExpectations(t)
}
