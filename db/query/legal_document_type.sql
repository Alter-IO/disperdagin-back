-- name: FindLegalDocTypeByID :one
SELECT
    id,
    description,
    author,
    created_at,
    updated_at
FROM
    legal_document_types
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: FindAllLegalDocTypes :many
SELECT
    id,
    description,
    author,
    created_at
FROM
    legal_document_types
WHERE
    deleted_at IS NULL
ORDER BY
    description ASC;

-- name: InsertLegalDocType :exec
INSERT INTO legal_document_types(id, description, author, created_at)
VALUES (
    sqlc.arg(id),
    sqlc.arg(description),
    sqlc.arg(author),
    sqlc.arg(created_at)
);

-- name: UpdateLegalDocType :execrows
UPDATE
    legal_document_types
SET
    description = sqlc.arg(description),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteLegalDocType :execrows
UPDATE
    legal_document_types
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;