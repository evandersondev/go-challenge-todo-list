package repository

import (
	"database/sql"

	"github.com/evandersondev/test-golang-todo-list/internal/dto"
	"github.com/evandersondev/test-golang-todo-list/internal/entity"
)

type TodoRepositoryInterface interface {
	Create(todo entity.Todo) error
	FindAll() ([]entity.Todo, error)
	FindById(id int) (*entity.Todo, error)
	Update(id int, dto dto.UpdateTodoDTO) (*entity.Todo, error)
	Delete(id int) error
}

var _ TodoRepositoryInterface = &TodoRepository{}

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		DB: db,
	}
}

func (r *TodoRepository) Create(todo entity.Todo) error {
	_, err := r.DB.Exec("INSERT INTO todos (title, description, done, created_at) VALUES (?, ?, ?, ?)", todo.Title, todo.Description, todo.Done, todo.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) FindAll() ([]entity.Todo, error) {
	rows, err := r.DB.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []entity.Todo
	for rows.Next() {
		var todo entity.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done, &todo.CreatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *TodoRepository) FindById(id int) (*entity.Todo, error) {
	var todo entity.Todo
	err := r.DB.QueryRow("SELECT id, title, description, done, created_at FROM todos WHERE id = ?", id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done, &todo.CreatedAt)
	if err != nil {
		return &entity.Todo{}, err
	}
	return &todo, nil
}

func (r *TodoRepository) Update(id int, dto dto.UpdateTodoDTO) (*entity.Todo, error) {
	rows, err := r.DB.Exec("UPDATE todos SET title = ?, description = ?, done = ? WHERE id = ?", dto.Title, dto.Description, dto.Done, id)
	if err != nil {
		return &entity.Todo{}, err
	}
	affectedRows, err := rows.RowsAffected()
	if err != nil || affectedRows == 0 {
		return &entity.Todo{}, err
	}
	todo, err := r.FindById(id)
	if err != nil {
		return &entity.Todo{}, err
	}

	return todo, nil
}

func (r *TodoRepository) Delete(id int) error {
	err := r.DB.QueryRow("DELETE FROM todos WHERE id = ?", id).Err()
	if err != nil {
		return err
	}
	return nil
}
