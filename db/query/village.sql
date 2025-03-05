-- name: FindVillageByID :one
SELECT
    id,
    subdistrict_id,
    name,
    author,
    created_at,
    updated_at
FROM
    villages
WHERE
    id = $1
AND
    deleted_at IS NULL;

-- name: FindAllVillages :many
SELECT
    id,
    subdistrict_id,
    name,
    author,
    created_at
FROM
    villages
WHERE
    deleted_at IS NULL
ORDER BY
    name ASC;

-- name: FindVillagesBySubdistrict :many
SELECT
    id,
    subdistrict_id,
    name,
    author,
    created_at
FROM
    villages
WHERE
    subdistrict_id = $1
AND
    deleted_at IS NULL
ORDER BY
    name ASC;

-- name: InsertVillage :exec
INSERT INTO villages(id, subdistrict_id, name, author, created_at)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateVillage :execrows
UPDATE
    villages
SET
    subdistrict_id = sqlc.arg(subdistrict_id),
    name = sqlc.arg(name),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteVillage :execrows
UPDATE
    villages
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;