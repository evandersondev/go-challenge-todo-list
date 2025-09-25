package handler

import (
	"encoding/json"
	"net/http"

	"github.com/evandersondev/test-golang-todo-list/internal/dto"
	"github.com/evandersondev/test-golang-todo-list/internal/usecase"
)

type TodoHandler struct {
	createTodoUsecase *usecase.CreateTodoUseCase
	findAllUseCase    *usecase.FindAllUseCase
	findTodoById      *usecase.FindTodoById
	updateTodo        *usecase.UpdateTodo
	deleteTodo        *usecase.DeleteTodo
}

func NewTodoHandler(
	createTodoUsecase *usecase.CreateTodoUseCase,
	findAllUseCase *usecase.FindAllUseCase,
	findTodoById *usecase.FindTodoById,
	updateTodo *usecase.UpdateTodo,
	deleteTodo *usecase.DeleteTodo,
) *TodoHandler {
	return &TodoHandler{
		createTodoUsecase: createTodoUsecase,
		findAllUseCase:    findAllUseCase,
		findTodoById:      findTodoById,
		updateTodo:        updateTodo,
		deleteTodo:        deleteTodo,
	}
}

func (h *TodoHandler) CreateTodoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		defer body.Close()
		dto := dto.CreateTodoDTO{}
		err := json.NewDecoder(body).Decode(&dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = h.createTodoUsecase.Execute(dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *TodoHandler) FindAllTodosHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := h.findAllUseCase.Execute()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	}
}

func (h *TodoHandler) FindTodoByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "ID parameter is missing", http.StatusBadRequest)
			return
		}
		todo, err := h.findTodoById.Execute(id)
		if err != nil {
			http.Error(w, "Todo not found", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todo)
	}
}

func (h *TodoHandler) UpdateTodoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "ID parameter is missing", http.StatusBadRequest)
			return
		}
		body := r.Body
		defer body.Close()
		dto := dto.UpdateTodoDTO{}
		err := json.NewDecoder(body).Decode(&dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		todo, err := h.updateTodo.Execute(id, dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todo)
	}
}

func (h *TodoHandler) DeleteTodoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "ID parameter is missing", http.StatusBadRequest)
			return
		}
		err := h.deleteTodo.Execute(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	}
}
