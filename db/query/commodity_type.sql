-- name: FindCommodityTypeByID :one
SELECT
    id,
    description,
    author,
    created_at,
    updated_at
FROM
    commodity_types
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: FindAllCommodityTypes :many
SELECT
    id,
    description,
    author,
    created_at
FROM
    commodity_types
WHERE
    deleted_at IS NULL
ORDER BY
    description ASC;

-- name: InsertCommodityType :exec
INSERT INTO commodity_types(id, description, author, created_at)
VALUES (
    sqlc.arg(id),
    sqlc.arg(description),
    sqlc.arg(author),
    sqlc.arg(created_at)
);

-- name: UpdateCommodityType :execrows
UPDATE
    commodity_types
SET
    description = sqlc.arg(description),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteCommodityType :execrows
UPDATE
    commodity_types
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;