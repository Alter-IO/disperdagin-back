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
    id = sqlc.arg(id)
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
    subdistrict_id = sqlc.arg(subdistrict_id)
AND
    deleted_at IS NULL
ORDER BY
    name ASC;

-- name: InsertVillage :exec
INSERT INTO villages(id, subdistrict_id, name, author, created_at)
VALUES (
    sqlc.arg(id),
    sqlc.arg(subdistrict_id),
    sqlc.arg(name),
    sqlc.arg(author),
    sqlc.arg(created_at)
);

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