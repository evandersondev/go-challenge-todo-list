package repository

import (
	"context"

	"github.com/evandersondev/test-golang-todo-list/internal/db"
	"github.com/evandersondev/test-golang-todo-list/internal/dto"
	"github.com/evandersondev/test-golang-todo-list/internal/entity"
	"github.com/google/uuid"
)

type TodoRepositoryInterface interface {
	Create(todo entity.Todo) error
	FindAll() ([]entity.Todo, error)
	FindById(id string) (*entity.Todo, error)
	Update(id string, dto dto.UpdateTodoDTO) (*entity.Todo, error)
	Delete(id string) error
}

var _ TodoRepositoryInterface = &TodoRepository{}

type TodoRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewTodoRepository(ctx context.Context, db *db.Queries) *TodoRepository {
	return &TodoRepository{
		ctx: ctx,
		db:  db,
	}
}

func (r *TodoRepository) Create(todo entity.Todo) error {
	err := r.db.CreateTodo(r.ctx, db.CreateTodoParams{
		ID:          uuid.New().String(),
		Title:       todo.Title,
		Description: todo.Description,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) FindAll() ([]entity.Todo, error) {
	rows, err := r.db.ListTodos(r.ctx)
	if err != nil {
		return nil, err
	}
	todos := make([]entity.Todo, 0, len(rows))
	for _, row := range rows {
		todos = append(todos, entity.Todo{
			ID:          string(row.ID),
			Title:       row.Title,
			Description: row.Description,
			Done:        row.Done,
			CreatedAt:   row.CreatedAt,
		})
	}
	return todos, nil
}

func (r *TodoRepository) FindById(id string) (*entity.Todo, error) {
	row, err := r.db.GetTodo(r.ctx, id)
	if err != nil {
		return &entity.Todo{}, err
	}
	return &entity.Todo{
		ID:          string(row.ID),
		Title:       row.Title,
		Description: row.Description,
		Done:        row.Done,
		CreatedAt:   row.CreatedAt,
	}, nil
}

func (r *TodoRepository) Update(id string, dto dto.UpdateTodoDTO) (*entity.Todo, error) {
	err := r.db.UpdateTodo(r.ctx, db.UpdateTodoParams{
		ID:          id,
		Title:       dto.Title,
		Description: dto.Description,
		Done:        dto.Done,
	})
	if err != nil {
		return &entity.Todo{}, err
	}
	todo, err := r.FindById(id)
	if err != nil {
		return &entity.Todo{}, err
	}
	return todo, nil
}

func (r *TodoRepository) Delete(id string) error {
	err := r.db.DeleteTodo(r.ctx, id)
	if err != nil {
		return err
	}
	return nil
}
