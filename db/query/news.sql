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
    id = $1
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
VALUES ($1, $2, $3, $4, $5);

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