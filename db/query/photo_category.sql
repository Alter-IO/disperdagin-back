-- name: FindPhotoCategoryByID :one
SELECT
    id,
    category,
    author,
    created_at,
    updated_at
FROM
    photo_categories
WHERE
    id = $1
AND
    deleted_at IS NULL;

-- name: FindAllPhotoCategories :many
SELECT
    id,
    category,
    author,
    created_at
FROM
    photo_categories
WHERE
    deleted_at IS NULL
ORDER BY
    category ASC;

-- name: InsertPhotoCategory :exec
INSERT INTO photo_categories(id, category, author, created_at)
VALUES ($1, $2, $3, $4);

-- name: UpdatePhotoCategory :execrows
UPDATE
    photo_categories
SET
    category = sqlc.arg(category),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeletePhotoCategory :execrows
UPDATE
    photo_categories
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;