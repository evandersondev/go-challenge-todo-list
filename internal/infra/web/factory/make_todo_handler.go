package factory

import (
	"database/sql"

	"github.com/evandersondev/test-golang-todo-list/internal/infra/repository"
	"github.com/evandersondev/test-golang-todo-list/internal/infra/web/handler"
	"github.com/evandersondev/test-golang-todo-list/internal/usecase"
)

func MakeTodoHandler(db *sql.DB) *handler.TodoHandler {
	repository := repository.NewTodoRepository(db)

	createTodoUsecase := usecase.NewCreateTodoUseCase(repository)
	findAllUseCase := usecase.NewFindAllUseCase(repository)
	findTodoById := usecase.NewFindTodoById(repository)
	updateTodo := usecase.NewUpdateTodo(repository)
	deleteTodo := usecase.NewDeleteTodo(repository)

	return handler.NewTodoHandler(createTodoUsecase, findAllUseCase, findTodoById, updateTodo, deleteTodo)
}
