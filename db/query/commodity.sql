-- name: FindCommodityByID :one
SELECT
    c.id,
    c.commodity_type_id,
    ct.description as commodity_type_name,
    c.name,
    c.unit,
    c.description,
    c.author,
    c.created_at,
    c.updated_at
FROM
    commodities c
LEFT JOIN
    commodity_types ct ON c.commodity_type_id = ct.id AND ct.deleted_at IS NULL
WHERE
    c.id = sqlc.arg(id)
AND
    c.deleted_at IS NULL;

-- name: FindAllCommodities :many
SELECT
    id,
    name,
    unit,
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
    unit,
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

-- name: InsertCommodity :exec
INSERT INTO commodities(id, name, unit, description, commodity_type_id, author, created_at)
VALUES (
    sqlc.arg(id),
    sqlc.arg(name),
    sqlc.arg(unit),
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
    unit = sqlc.arg(unit),
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

-- name: FindLatestCommodities :one
SELECT
    id,
    commodities,
    publish_date
FROM
    daily_commodities
ORDER BY
    publish_date DESC
LIMIT 1;

-- name: InsertDailyCommodity :exec
INSERT INTO daily_commodities (id, commodities, publish_date)
VALUES (sqlc.arg(id), sqlc.arg(commodities), sqlc.arg(publish_date));
