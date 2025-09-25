-- name: ListTodos :many
SELECT * FROM todos;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = ?;

-- name: CreateTodo :exec
INSERT INTO todos (id, title, description)
VALUES (?, ?, ?);

-- name: UpdateTodo :exec
UPDATE todos SET title = ?, description = ?, done = ? WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = ?;