-- name: FindUserByID :one
SELECT
    id,
    role_id,
    username,
    password
FROM
    users
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: FindUsers :many
SELECT
    id,
    role_id,
    username
FROM
    users
WHERE
    deleted_at IS NULL
ORDER BY
    id DESC;

-- name: InsertUser :exec
INSERT INTO users(id, role_id, username, password, created_at)
VALUES (sqlc.arg(id), sqlc.arg(role_id), sqlc.arg(username), sqlc.arg(password), sqlc.arg(created_at));

-- name: UpdatePassword :execrows
UPDATE
    users
SET
    password = sqlc.arg(password),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteUser :execrows
UPDATE
    users
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;