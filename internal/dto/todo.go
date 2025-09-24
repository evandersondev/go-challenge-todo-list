package dto

type CreateTodoDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTodoDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
