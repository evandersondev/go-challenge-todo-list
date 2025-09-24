package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/evandersondev/test-golang-todo-list/internal/infra/web/factory"
	_ "github.com/mattn/go-sqlite3"
)

// db.Exec("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, description TEXT, done BOOLEAN, created_at DATETIME)")

func main() {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	todoHanlder := factory.MakeTodoHandler(db)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /todos", todoHanlder.CreateTodoHandler())
	mux.HandleFunc("GET /todos", todoHanlder.FindAllTodosHandler())
	mux.HandleFunc("GET /todos/{id}", todoHanlder.FindTodoByIdHandler())
	mux.HandleFunc("PUT /todos/{id}", todoHanlder.UpdateTodoHandler())
	mux.HandleFunc("DELETE /todos/{id}", todoHanlder.DeleteTodoHandler())

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}
