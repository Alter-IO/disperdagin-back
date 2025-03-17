-- name: FindLegalDocumentByID :one
SELECT
    id,
    document_name,
    file_name,
    document_type,
    description,
    author,
    created_at,
    updated_at
FROM
    legal_documents
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: FindAllLegalDocuments :many
SELECT
    id,
    document_name,
    file_name,
    document_type,
    author,
    created_at
FROM
    legal_documents
WHERE
    deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: FindLegalDocumentsByType :many
SELECT
    id,
    document_name,
    file_name,
    document_type,
    author,
    created_at
FROM
    legal_documents
WHERE
    document_type = sqlc.arg(document_type)
AND
    deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: InsertLegalDocument :exec
INSERT INTO legal_documents(id, document_name, file_name, document_type, description, author, created_at)
VALUES (
    sqlc.arg(id),
    sqlc.arg(document_name),
    sqlc.arg(file_name),
    sqlc.arg(document_type),
    sqlc.arg(description),
    sqlc.arg(author),
    sqlc.arg(created_at)
);

-- name: UpdateLegalDocument :execrows
UPDATE
    legal_documents
SET
    document_name = sqlc.arg(document_name),
    file_name = sqlc.arg(file_name),
    document_type = sqlc.arg(document_type),
    description = sqlc.arg(description),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteLegalDocument :execrows
UPDATE
    legal_documents
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;