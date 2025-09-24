package entity

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewTodo(title, description string) *Todo {
	return &Todo{
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
	}
}

func (t *Todo) ToogleDone() {
	t.Done = !t.Done
}

func (t *Todo) IsValid() bool {
	if t.Title == "" {
		return false
	}

	if len(t.Description) > 255 {
		return false
	}

	return true
}
