-- name: IKM :one
SELECT
    id,
    description,
    village_id,
    business_type,
    author,
    created_at,
    updated_at
FROM
    ikms
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: FindAllIKMs :many
SELECT
    id,
    description,
    village_id,
    business_type,
    author,
    created_at
FROM
    ikms
WHERE
    deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: FindIKMsByVillage :many
SELECT
    id,
    description,
    village_id,
    business_type,
    author,
    created_at
FROM
    ikms
WHERE
    village_id = sqlc.arg(village_id)
AND
    deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: FindIKMsByBusinessType :many
SELECT
    id,
    description,
    village_id,
    business_type,
    author,
    created_at
FROM
    ikms
WHERE
    business_type = sqlc.arg(business_type)
AND
    deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: InsertIKM :exec
INSERT INTO ikms(id, description, village_id, business_type, author, created_at)
VALUES (
    sqlc.arg(id),
    sqlc.arg(description),
    sqlc.arg(village_id),
    sqlc.arg(business_type),
    sqlc.arg(author),
    sqlc.arg(created_at)
);

-- name: UpdateIKM :execrows
UPDATE
    ikms
SET
    description = sqlc.arg(description),
    village_id = sqlc.arg(village_id),
    business_type = sqlc.arg(business_type),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteIKM :execrows
UPDATE
    ikms
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;