-- name: FindMarketByID :one
SELECT
    id,
    name,
    author,
    created_at,
    updated_at
FROM
    markets
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: FindAllMarkets :many
SELECT
    id,
    name,
    author,
    created_at
FROM
    markets
WHERE
    deleted_at IS NULL
ORDER BY
    name ASC;

-- name: InsertMarket :exec
INSERT INTO markets(id, name, author, created_at)
VALUES (
    sqlc.arg(id),
    sqlc.arg(name),
    sqlc.arg(author),
    sqlc.arg(created_at)
);

-- name: UpdateMarket :execrows
UPDATE
    markets
SET
    name = sqlc.arg(name),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteMarket :execrows
UPDATE
    markets
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;