-- name: FindPublicInfoTypeByID :one
SELECT
    id,
    description,
    author,
    created_at,
    updated_at
FROM
    public_information_types
WHERE
    id = $1
AND
    deleted_at IS NULL;

-- name: FindAllPublicInfoTypes :many
SELECT
    id,
    description,
    author,
    created_at
FROM
    public_information_types
WHERE
    deleted_at IS NULL
ORDER BY
    description ASC;

-- name: InsertPublicInfoType :exec
INSERT INTO public_information_types(id, description, author, created_at)
VALUES ($1, $2, $3, $4);

-- name: UpdatePublicInfoType :execrows
UPDATE
    public_information_types
SET
    description = sqlc.arg(description),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeletePublicInfoType :execrows
UPDATE
    public_information_types
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;