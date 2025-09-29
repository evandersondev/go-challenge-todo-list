package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewTodo(title, description string) (*Todo, error) {
	todo := Todo{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
	}

	if !todo.isValid() {
		return nil, errors.New("invalid todo")
	}

	return &todo, nil
}

func (t *Todo) ToogleDone() {
	t.Done = !t.Done
}

func (t *Todo) isValid() bool {
	if t.Title == "" || len(t.Description) < 20 {
		return false
	}

	return true
}
