-- name: FindMarketFeeByID :one
SELECT
    id,
    market_id,
    num_permanent_kiosks,
    num_non_permanent_kiosks,
    permanent_kiosk_revenue,
    non_permanent_kiosk_revenue,
    collection_status,
    description,
    semester,
    year,
    author,
    created_at,
    updated_at
FROM
    market_fees
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: FindAllMarketFees :many
SELECT
    id,
    market_id,
    permanent_kiosk_revenue,
    non_permanent_kiosk_revenue,
    collection_status,
    semester,
    year,
    author,
    created_at
FROM
    market_fees
WHERE
    deleted_at IS NULL
ORDER BY
    year DESC, semester DESC;

-- name: FindMarketFeesByMarket :many
SELECT
    id,
    market_id,
    permanent_kiosk_revenue,
    non_permanent_kiosk_revenue,
    collection_status,
    semester,
    year,
    author,
    created_at
FROM
    market_fees
WHERE
    market_id = sqlc.arg(market_id)
AND
    deleted_at IS NULL
ORDER BY
    year DESC, semester DESC;

-- name: FindMarketFeesByYear :many
SELECT
    id,
    market_id,
    permanent_kiosk_revenue,
    non_permanent_kiosk_revenue,
    collection_status,
    semester,
    year,
    author,
    created_at
FROM
    market_fees
WHERE
    year = sqlc.arg(year)
AND
    deleted_at IS NULL
ORDER BY
    market_id, semester;

-- name: FindMarketFeesBySemesterAndYear :many
SELECT
    id,
    market_id,
    permanent_kiosk_revenue,
    non_permanent_kiosk_revenue,
    collection_status,
    semester,
    year,
    author,
    created_at
FROM
    market_fees
WHERE
    semester = sqlc.arg(semester)
AND
    year = sqlc.arg(year)
AND
    deleted_at IS NULL
ORDER BY
    market_id;

-- name: InsertMarketFee :exec
INSERT INTO market_fees(
    id, 
    market_id, 
    num_permanent_kiosks, 
    num_non_permanent_kiosks, 
    permanent_kiosk_revenue, 
    non_permanent_kiosk_revenue, 
    collection_status, 
    description, 
    semester, 
    year, 
    author, 
    created_at
)
VALUES (
    sqlc.arg(id), 
    sqlc.arg(market_id), 
    sqlc.arg(num_permanent_kiosks), 
    sqlc.arg(num_non_permanent_kiosks), 
    sqlc.arg(permanent_kiosk_revenue), 
    sqlc.arg(non_permanent_kiosk_revenue), 
    sqlc.arg(collection_status), 
    sqlc.arg(description), 
    sqlc.arg(semester), 
    sqlc.arg(year), 
    sqlc.arg(author), 
    sqlc.arg(created_at)
);

-- name: UpdateMarketFee :execrows
UPDATE
    market_fees
SET
    market_id = sqlc.arg(market_id),
    num_permanent_kiosks = sqlc.arg(num_permanent_kiosks),
    num_non_permanent_kiosks = sqlc.arg(num_non_permanent_kiosks),
    permanent_kiosk_revenue = sqlc.arg(permanent_kiosk_revenue),
    non_permanent_kiosk_revenue = sqlc.arg(non_permanent_kiosk_revenue),
    collection_status = sqlc.arg(collection_status),
    description = sqlc.arg(description),
    semester = sqlc.arg(semester),
    year = sqlc.arg(year),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteMarketFee :execrows
UPDATE
    market_fees
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;