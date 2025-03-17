-- name: FindGreetingByID :one
SELECT
    id,
    message,
    author,
    created_at,
    updated_at
FROM
    greetings
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: FindLatestGreeting :one
SELECT
    id,
    message,
    author,
    created_at,
    updated_at
FROM
    greetings
WHERE
    deleted_at IS NULL
ORDER BY
    created_at DESC
LIMIT 1;

-- name: FindAllGreetings :many
SELECT
    id,
    message,
    author,
    created_at
FROM
    greetings
WHERE
    deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: InsertGreeting :exec
INSERT INTO greetings(id, message, author, created_at)
VALUES (
    sqlc.arg(id),
    sqlc.arg(message),
    sqlc.arg(author),
    sqlc.arg(created_at)
);

-- name: UpdateGreeting :execrows
UPDATE
    greetings
SET
    message = sqlc.arg(message),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteGreeting :execrows
UPDATE
    greetings
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;