-- name: FindVideoByID :one
SELECT
    id,
    title,
    link,
    description,
    author,
    created_at,
    updated_at
FROM
    videos
WHERE
    id = $1
AND
    deleted_at IS NULL;

-- name: FindAllVideos :many
SELECT
    id,
    title,
    link,
    author,
    created_at
FROM
    videos
WHERE
    deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: InsertVideo :exec
INSERT INTO videos(id, title, link, description, author, created_at)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateVideo :execrows
UPDATE
    videos
SET
    title = sqlc.arg(title),
    link = sqlc.arg(link),
    description = sqlc.arg(description),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteVideo :execrows
UPDATE
    videos
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;