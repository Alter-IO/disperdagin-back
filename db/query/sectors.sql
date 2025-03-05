-- name: FindSectorByID :one
SELECT
    id,
    name,
    description,
    author,
    created_at,
    updated_at
FROM
    sectors
WHERE
    id = $1
AND
    deleted_at IS NULL;

-- name: FindAllSectors :many
SELECT
    id,
    name,
    author,
    created_at
FROM
    sectors
WHERE
    deleted_at IS NULL
ORDER BY
    name ASC;

-- name: InsertSector :exec
INSERT INTO sectors(id, name, description, author, created_at)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateSector :execrows
UPDATE
    sectors
SET
    name = sqlc.arg(name),
    description = sqlc.arg(description),
    author = sqlc.arg(author),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteSector :execrows
UPDATE
    sectors
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;