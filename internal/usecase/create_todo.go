package usecase

import (
	"github.com/evandersondev/test-golang-todo-list/internal/dto"
	"github.com/evandersondev/test-golang-todo-list/internal/entity"
	"github.com/evandersondev/test-golang-todo-list/internal/infra/repository"
)

type CreateTodoUseCase struct {
	repository repository.TodoRepositoryInterface
}

func NewCreateTodoUseCase(repository repository.TodoRepositoryInterface) *CreateTodoUseCase {
	return &CreateTodoUseCase{
		repository: repository,
	}
}

func (uc *CreateTodoUseCase) Execute(dto dto.CreateTodoDTO) (*entity.Todo, error) {
	todo := entity.NewTodo(dto.Title, dto.Description)

	if !todo.IsValid() {
		return &entity.Todo{}, nil
	}

	err := uc.repository.Create(*todo)
	if err != nil {
		return &entity.Todo{}, err
	}

	return todo, nil
}
