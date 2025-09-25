package usecase

import (
	"github.com/evandersondev/test-golang-todo-list/internal/entity"
	"github.com/evandersondev/test-golang-todo-list/internal/infra/repository"
)

type FindTodoById struct {
	repository repository.TodoRepositoryInterface
}

func NewFindTodoById(repository repository.TodoRepositoryInterface) *FindTodoById {
	return &FindTodoById{
		repository: repository,
	}
}

func (uc *FindTodoById) Execute(id string) (*entity.Todo, error) {
	todo, err := uc.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
