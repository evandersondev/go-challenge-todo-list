package usecase

import (
	"github.com/evandersondev/test-golang-todo-list/internal/entity"
	"github.com/evandersondev/test-golang-todo-list/internal/infra/repository"
)

type FindAllUseCase struct {
	repository repository.TodoRepositoryInterface
}

func NewFindAllUseCase(repository repository.TodoRepositoryInterface) *FindAllUseCase {
	return &FindAllUseCase{
		repository: repository,
	}
}

func (uc *FindAllUseCase) Execute() ([]entity.Todo, error) {
	todos, err := uc.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}
