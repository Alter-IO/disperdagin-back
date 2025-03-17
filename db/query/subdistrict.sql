-- name: FindSubdistrictByID :one
SELECT
    id,
    name,
    author,
    created_at,
    updated_at
FROM
    subdistricts
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: FindAllSubdistricts :many
SELECT
    id,
    name,
    author,
    created_at
FROM
    subdistricts
WHERE
    deleted_at IS NULL
ORDER BY
    name ASC;

-- name: InsertSubdistrict :exec
INSERT INTO subdistricts(id, name, author, created_at)
VALUES (
    sqlc.arg(id),
    sqlc.arg(name),
    sqlc.arg(author),
    sqlc.arg(created_at)
);

-- name: UpdateSubdistrict :execrows
UPDATE
    subdistricts
SET
    name = sqlc.arg(name),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteSubdistrict :execrows
UPDATE
    subdistricts
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;