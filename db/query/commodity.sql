-- name: FindCommodityByID :one
SELECT
    id,
    name,
    price,
    unit,
    publish_date,
    description,
    commodity_type_id,
    author,
    created_at,
    updated_at
FROM
    commodities
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: FindAllCommodities :many
SELECT
    id,
    name,
    price,
    unit,
    publish_date,
    commodity_type_id,
    author,
    created_at
FROM
    commodities
WHERE
    deleted_at IS NULL
ORDER BY
    name ASC;

-- name: FindCommoditiesByType :many
SELECT
    id,
    name,
    price,
    unit,
    publish_date,
    commodity_type_id,
    author,
    created_at
FROM
    commodities
WHERE
    commodity_type_id = sqlc.arg(commodity_type_id)
AND
    deleted_at IS NULL
ORDER BY
    name ASC;

-- name: FindLatestCommodities :many
SELECT
    id,
    name,
    price,
    unit,
    publish_date,
    commodity_type_id,
    author,
    created_at
FROM
    commodities
WHERE
    deleted_at IS NULL
ORDER BY
    publish_date DESC;

-- name: InsertCommodity :exec
INSERT INTO commodities(id, name, price, unit, publish_date, description, commodity_type_id, author, created_at)
VALUES (
    sqlc.arg(id),
    sqlc.arg(name),
    sqlc.arg(price),
    sqlc.arg(unit),
    sqlc.arg(publish_date),
    sqlc.arg(description),
    sqlc.arg(commodity_type_id),
    sqlc.arg(author),
    sqlc.arg(created_at)
);

-- name: UpdateCommodity :execrows
UPDATE
    commodities
SET
    name = sqlc.arg(name),
    price = sqlc.arg(price),
    unit = sqlc.arg(unit),
    publish_date = sqlc.arg(publish_date),
    description = sqlc.arg(description),
    commodity_type_id = sqlc.arg(commodity_type_id),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteCommodity :execrows
UPDATE
    commodities
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;