package usecase

import "github.com/evandersondev/test-golang-todo-list/internal/infra/repository"

type DeleteTodo struct {
	repository repository.TodoRepositoryInterface
}

func NewDeleteTodo(repository *repository.TodoRepository) *DeleteTodo {
	return &DeleteTodo{
		repository: repository,
	}
}

func (uc *DeleteTodo) Execute(id string) error {
	err := uc.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
