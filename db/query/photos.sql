-- name: FindPhotoByID :one
SELECT
    id,
    category_id,
    title,
    file_url,
    description,
    author,
    created_at,
    updated_at
FROM
    photos
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: FindAllPhotos :many
SELECT
    id,
    category_id,
    title,
    file_url,
    author,
    created_at
FROM
    photos
WHERE
    deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: FindPhotosByCategory :many
SELECT
    id,
    category_id,
    title,
    file_url,
    author,
    created_at
FROM
    photos
WHERE
    category_id = sqlc.arg(category_id)
AND
    deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: InsertPhoto :exec
INSERT INTO photos(id, category_id, title, file_url, description, author, created_at)
VALUES (
    sqlc.arg(id),
    sqlc.arg(category_id),
    sqlc.arg(title),
    sqlc.arg(file_url),
    sqlc.arg(description),
    sqlc.arg(author),
    sqlc.arg(created_at)
);

-- name: UpdatePhoto :execrows
UPDATE
    photos
SET
    category_id = sqlc.arg(category_id),
    title = sqlc.arg(title),
    description = sqlc.arg(description),
    file_url = sqlc.arg(file_url),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeletePhoto :execrows
UPDATE
    photos
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;