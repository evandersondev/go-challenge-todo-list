package usecase

import (
	"testing"

	"github.com/evandersondev/test-golang-todo-list/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewCreateTodoUseCase(t *testing.T) {
	repository := &TodoRepositoryMock{}
	sut := NewCreateTodoUseCase(repository)
	assert.IsTypef(t, &CreateTodoUseCase{}, sut, "Expected sut to be of type *CreateTodoUseCase")
}

func TestCreateTodoUseCase_Execute_Success(t *testing.T) {
	repository := &TodoRepositoryMock{}
	repository.On("Create", mock.Anything).Return(nil).Once()
	sut := NewCreateTodoUseCase(repository)

	dto := dto.CreateTodoDTO{
		Title:       "Test Todo",
		Description: "This is a test description for the todo",
	}

	todo, err := sut.Execute(dto)

	assert.NoError(t, err)
	assert.NotNil(t, todo)
	assert.Equal(t, dto.Title, todo.Title)
	assert.Equal(t, dto.Description, todo.Description)
	assert.False(t, todo.Done)
	assert.NotEmpty(t, todo.ID)
	assert.NotEmpty(t, todo.CreatedAt)

	repository.AssertExpectations(t)
}

func TestCreateTodoUseCase_Execute_Failure(t *testing.T) {
	repository := &TodoRepositoryMock{}
	repository.On("Create", mock.Anything).Return(assert.AnError)
	sut := NewCreateTodoUseCase(repository)

	dto := dto.CreateTodoDTO{
		Title:       "Test Todo",
		Description: "This is a test description for the todo",
	}
	_, err := sut.Execute(dto)

	assert.Error(t, err)
	repository.AssertExpectations(t)
}

func TestCreateTodoUseCase_Execute_InvalidTodo(t *testing.T) {
	repository := &TodoRepositoryMock{}
	sut := NewCreateTodoUseCase(repository)

	dto := dto.CreateTodoDTO{
		Title:       "Test Todo",
		Description: "Short description",
	}
	_, err := sut.Execute(dto)

	assert.Error(t, err)
	assert.EqualError(t, err, "invalid todo")
	repository.AssertExpectations(t)
}
