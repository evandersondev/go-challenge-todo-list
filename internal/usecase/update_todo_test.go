package usecase

import (
	"testing"

	"github.com/evandersondev/test-golang-todo-list/internal/dto"
	"github.com/evandersondev/test-golang-todo-list/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewUpdateTodoUseCase(t *testing.T) {
	repository := &TodoRepositoryMock{}
	sut := NewUpdateTodoUseCase(repository)
	assert.IsTypef(t, &UpdateTodoUseCase{}, sut, "Expected sut to be of type *UpdateTodoUseCase")
}

func TestUpdateTodoUseCase_Execute_Success(t *testing.T) {
	repository := &TodoRepositoryMock{}
	repository.On("Update", "1", mock.Anything).Return(&entity.Todo{
		ID:          "1",
		Title:       "Updated Title",
		Description: "Updated Description",
		Done:        true,
	}, nil)
	sut := NewUpdateTodoUseCase(repository)

	todo, err := sut.Execute("1", dto.UpdateTodoDTO{
		Title:       "Updated Title",
		Description: "Updated Description",
		Done:        true,
	})
	assert.NoError(t, err)
	assert.NotNil(t, todo)
	assert.Equal(t, "Updated Title", todo.Title)
	assert.Equal(t, "Updated Description", todo.Description)
	assert.True(t, todo.Done)
	repository.AssertExpectations(t)
}

func TestUpdateTodoUseCase_Execute_Failure(t *testing.T) {
	repository := &TodoRepositoryMock{}
	repository.On("Update", "1", mock.Anything).Return(nil, assert.AnError)
	sut := NewUpdateTodoUseCase(repository)

	todo, err := sut.Execute("1", dto.UpdateTodoDTO{
		Title:       "Updated Title",
		Description: "Updated Description",
		Done:        true,
	})
	assert.Error(t, err)
	assert.Nil(t, todo)
	repository.AssertExpectations(t)
}
