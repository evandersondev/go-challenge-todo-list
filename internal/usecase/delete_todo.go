package usecase

import "github.com/evandersondev/test-golang-todo-list/internal/infra/repository"

type DeleteTodoUseCase struct {
	repository repository.TodoRepositoryInterface
}

func NewDeleteTodoUseCase(repository repository.TodoRepositoryInterface) *DeleteTodoUseCase {
	return &DeleteTodoUseCase{
		repository: repository,
	}
}

func (uc *DeleteTodoUseCase) Execute(id string) error {
	err := uc.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
