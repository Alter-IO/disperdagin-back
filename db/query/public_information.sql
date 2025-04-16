-- name: FindPublicInfoByID :one
SELECT
    id,
    document_name,
    file_url,
    public_info_type,
    description,
    author,
    created_at,
    updated_at
FROM
    public_information
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: FindAllPublicInfo :many
SELECT
    id,
    document_name,
    file_url,
    public_info_type,
    author,
    created_at
FROM
    public_information
WHERE
    deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: FindPublicInfoByType :many
SELECT
    id,
    document_name,
    file_url,
    public_info_type,
    author,
    created_at
FROM
    public_information
WHERE
    public_info_type = sqlc.arg(public_info_type)
AND
    deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: InsertPublicInfo :exec
INSERT INTO public_information(id, document_name, file_url, public_info_type, description, author, created_at)
VALUES (
    sqlc.arg(id),
    sqlc.arg(document_name),
    sqlc.arg(file_url),
    sqlc.arg(public_info_type),
    sqlc.arg(description),
    sqlc.arg(author),
    sqlc.arg(created_at)
);

-- name: UpdatePublicInfo :execrows
UPDATE
    public_information
SET
    document_name = sqlc.arg(document_name),
    file_url = sqlc.arg(file_url),
    public_info_type = sqlc.arg(public_info_type),
    description = sqlc.arg(description),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeletePublicInfo :execrows
UPDATE
    public_information
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;