-- name: FindNewsById :one
SELECT
    id,
    title,
    content,
    author,
    created_at,
    updated_at
FROM
    news
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: FindAllNews :many
SELECT
    id,
    title,
    author,
    created_at
FROM
    news
WHERE
    deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: InsertNews :exec
INSERT INTO news(id, title, content, author, created_at)
VALUES (
    sqlc.arg(id),
    sqlc.arg(title),
    sqlc.arg(content),
    sqlc.arg(author),
    sqlc.arg(created_at)
);

-- name: UpdateNews :execrows
UPDATE
    news
SET
    title = sqlc.arg(title),
    content = sqlc.arg(content),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteNews :execrows
UPDATE
    news
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;