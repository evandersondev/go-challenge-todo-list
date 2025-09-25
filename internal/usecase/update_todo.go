package usecase

import (
	"github.com/evandersondev/test-golang-todo-list/internal/dto"
	"github.com/evandersondev/test-golang-todo-list/internal/entity"
	"github.com/evandersondev/test-golang-todo-list/internal/infra/repository"
)

type UpdateTodo struct {
	repository repository.TodoRepositoryInterface
}

func NewUpdateTodo(repository repository.TodoRepositoryInterface) *UpdateTodo {
	return &UpdateTodo{
		repository: repository,
	}
}

func (uc *UpdateTodo) Execute(id string, dto dto.UpdateTodoDTO) (*entity.Todo, error) {
	todo, err := uc.repository.Update(id, dto)
	if err != nil {
		return &entity.Todo{}, err
	}
	return todo, nil
}
