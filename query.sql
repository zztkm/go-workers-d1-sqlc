/* query for todo server */

-- name: GetTodo :one
SELECT *
FROM todos
WHERE id = ?
LIMIT 1;

-- name: ListTodos :many
SELECT *
FROM todos
ORDER BY created_at DESC;

-- name: CreateTodo :one
INSERT INTO todos (
	title
) VALUES (
	?
)
RETURNING *;

-- name: UpdateTodo :exec
UPDATE todos
SET title = ?
WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;

