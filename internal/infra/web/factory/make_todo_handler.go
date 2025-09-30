package factory

import (
	"context"

	"github.com/evandersondev/test-golang-todo-list/internal/db"
	"github.com/evandersondev/test-golang-todo-list/internal/infra/repository"
	"github.com/evandersondev/test-golang-todo-list/internal/infra/web/handler"
	"github.com/evandersondev/test-golang-todo-list/internal/usecase"
)

func MakeTodoHandler(ctx context.Context, db *db.Queries) *handler.TodoHandler {
	repository := repository.NewTodoRepository(ctx, db)

	createTodoUsecase := usecase.NewCreateTodoUseCase(repository)
	findAllUseCase := usecase.NewFindAllUseCase(repository)
	findTodoById := usecase.NewFindTodoById(repository)
	updateTodoUseCase := usecase.NewUpdateTodoUseCase(repository)
	deleteTodoUseCase := usecase.NewDeleteTodoUseCase(repository)

	return handler.NewTodoHandler(
		createTodoUsecase,
		findAllUseCase,
		findTodoById,
		updateTodoUseCase,
		deleteTodoUseCase,
	)
}
