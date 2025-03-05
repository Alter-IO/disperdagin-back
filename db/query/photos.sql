-- name: FindPhotoByID :one
SELECT
    id,
    category_id,
    title,
    file,
    description,
    author,
    created_at,
    updated_at
FROM
    photos
WHERE
    id = $1
AND
    deleted_at IS NULL;

-- name: FindAllPhotos :many
SELECT
    id,
    category_id,
    title,
    file,
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
    file,
    author,
    created_at
FROM
    photos
WHERE
    category_id = $1
AND
    deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: InsertPhoto :exec
INSERT INTO photos(id, category_id, title, file, description, author, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdatePhoto :execrows
UPDATE
    photos
SET
    category_id = sqlc.arg(category_id),
    title = sqlc.arg(title),
    description = sqlc.arg(description),
    file = sqlc.arg(file),
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