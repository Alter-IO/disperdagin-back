-- name: FindIKMTypeByID :one
SELECT
    id,
    document_name,
    file_name,
    public_info_type,
    description,
    author,
    created_at,
    updated_at
FROM
    ikm_types
WHERE
    id = $1
AND
    deleted_at IS NULL;

-- name: FindAllIKMTypes :many
SELECT
    id,
    document_name,
    file_name,
    public_info_type,
    author,
    created_at
FROM
    ikm_types
WHERE
    deleted_at IS NULL
ORDER BY
    document_name ASC;

-- name: FindIKMTypesByInfoType :many
SELECT
    id,
    document_name,
    file_name,
    public_info_type,
    author,
    created_at
FROM
    ikm_types
WHERE
    public_info_type = $1
AND
    deleted_at IS NULL
ORDER BY
    document_name ASC;

-- name: InsertIKMType :exec
INSERT INTO ikm_types(id, document_name, file_name, public_info_type, description, author, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateIKMType :execrows
UPDATE
    ikm_types
SET
    document_name = sqlc.arg(document_name),
    file_name = sqlc.arg(file_name),
    public_info_type = sqlc.arg(public_info_type),
    description = sqlc.arg(description),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteIKMType :execrows
UPDATE
    ikm_types
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;