package usecase

import (
	"github.com/evandersondev/test-golang-todo-list/internal/dto"
	"github.com/evandersondev/test-golang-todo-list/internal/entity"
	"github.com/evandersondev/test-golang-todo-list/internal/infra/repository"
)

type UpdateTodoUseCase struct {
	repository repository.TodoRepositoryInterface
}

func NewUpdateTodoUseCase(repository repository.TodoRepositoryInterface) *UpdateTodoUseCase {
	return &UpdateTodoUseCase{
		repository: repository,
	}
}

func (uc *UpdateTodoUseCase) Execute(id string, dto dto.UpdateTodoDTO) (*entity.Todo, error) {
	todo, err := uc.repository.Update(id, dto)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
