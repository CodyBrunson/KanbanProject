-- name: CreateNewTask :one
INSERT INTO tasks (id, title, description, status, created_at, updated_at)
VALUES (
gen_random_uuid(), $1, $2, $3, NOW(), NOW()
)
RETURNING *;

-- name: GetAllTasks :many
SELECT * FROM tasks;

-- name: GetTaskByID :one
SELECT * FROM tasks
WHERE id = $1;

-- name: UpdateTaskByID :exec
UPDATE tasks
SET title = $2, description = $3, updated_at = NOW()
WHERE id = $1;

-- name: FinishTask :exec
UPDATE tasks
SET status = 'FINISHED', updated_at = NOW(), completed_at = NOW()
WHERE id = $1;

-- name: DeleteTask :exec
UPDATE tasks
SET status = 'DELETED', updated_at = NOW(), deleted_at = NOW()
WHERE id = $1;