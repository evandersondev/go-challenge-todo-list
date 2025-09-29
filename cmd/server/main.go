package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/evandersondev/test-golang-todo-list/configs"
	"github.com/evandersondev/test-golang-todo-list/internal/db"
	"github.com/evandersondev/test-golang-todo-list/internal/infra/web/factory"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	dbc, err := sql.Open(config.DBDriver, config.DBUrl)
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
	http.ListenAndServe(":"+config.WebServerPort, mux)
}
