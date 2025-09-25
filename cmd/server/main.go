package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/evandersondev/test-golang-todo-list/internal/db"
	"github.com/evandersondev/test-golang-todo-list/internal/infra/web/factory"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbc, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/mydb?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer dbc.Close()
	queries := db.New(dbc)
	todoHanlder := factory.MakeTodoHandler(ctx, queries)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /todos", todoHanlder.CreateTodoHandler())
	mux.HandleFunc("GET /todos", todoHanlder.FindAllTodosHandler())
	mux.HandleFunc("GET /todos/{id}", todoHanlder.FindTodoByIdHandler())
	mux.HandleFunc("PUT /todos/{id}", todoHanlder.UpdateTodoHandler())
	mux.HandleFunc("DELETE /todos/{id}", todoHanlder.DeleteTodoHandler())

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}
