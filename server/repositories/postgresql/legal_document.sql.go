// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: legal_document.sql

package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteLegalDocument = `-- name: DeleteLegalDocument :execrows
UPDATE
    legal_documents
SET
    deleted_at = $1
WHERE
    id = $2
AND
    deleted_at IS NULL
`

type DeleteLegalDocumentParams struct {
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
	ID        pgtype.Text        `json:"id"`
}

func (q *Queries) DeleteLegalDocument(ctx context.Context, arg DeleteLegalDocumentParams) (int64, error) {
	result, err := q.db.Exec(ctx, deleteLegalDocument, arg.DeletedAt, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const findAllLegalDocuments = `-- name: FindAllLegalDocuments :many
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
    created_at DESC
`

type FindAllLegalDocumentsRow struct {
	ID           string             `json:"id"`
	DocumentName string             `json:"document_name"`
	FileName     string             `json:"file_name"`
	DocumentType string             `json:"document_type"`
	Author       string             `json:"author"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) FindAllLegalDocuments(ctx context.Context) ([]FindAllLegalDocumentsRow, error) {
	rows, err := q.db.Query(ctx, findAllLegalDocuments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FindAllLegalDocumentsRow{}
	for rows.Next() {
		var i FindAllLegalDocumentsRow
		if err := rows.Scan(
			&i.ID,
			&i.DocumentName,
			&i.FileName,
			&i.DocumentType,
			&i.Author,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findLegalDocumentByID = `-- name: FindLegalDocumentByID :one
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
    id = $1
AND
    deleted_at IS NULL
`

type FindLegalDocumentByIDRow struct {
	ID           string             `json:"id"`
	DocumentName string             `json:"document_name"`
	FileName     string             `json:"file_name"`
	DocumentType string             `json:"document_type"`
	Description  pgtype.Text        `json:"description"`
	Author       string             `json:"author"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) FindLegalDocumentByID(ctx context.Context, dollar_1 pgtype.Text) (FindLegalDocumentByIDRow, error) {
	row := q.db.QueryRow(ctx, findLegalDocumentByID, dollar_1)
	var i FindLegalDocumentByIDRow
	err := row.Scan(
		&i.ID,
		&i.DocumentName,
		&i.FileName,
		&i.DocumentType,
		&i.Description,
		&i.Author,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findLegalDocumentsByType = `-- name: FindLegalDocumentsByType :many
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
    document_type = $1
AND
    deleted_at IS NULL
ORDER BY
    created_at DESC
`

type FindLegalDocumentsByTypeRow struct {
	ID           string             `json:"id"`
	DocumentName string             `json:"document_name"`
	FileName     string             `json:"file_name"`
	DocumentType string             `json:"document_type"`
	Author       string             `json:"author"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) FindLegalDocumentsByType(ctx context.Context, dollar_1 pgtype.Text) ([]FindLegalDocumentsByTypeRow, error) {
	rows, err := q.db.Query(ctx, findLegalDocumentsByType, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FindLegalDocumentsByTypeRow{}
	for rows.Next() {
		var i FindLegalDocumentsByTypeRow
		if err := rows.Scan(
			&i.ID,
			&i.DocumentName,
			&i.FileName,
			&i.DocumentType,
			&i.Author,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertLegalDocument = `-- name: InsertLegalDocument :exec
INSERT INTO legal_documents(id, document_name, file_name, document_type, description, author, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type InsertLegalDocumentParams struct {
	Column1 pgtype.Text        `json:"column_1"`
	Column2 pgtype.Text        `json:"column_2"`
	Column3 pgtype.Text        `json:"column_3"`
	Column4 pgtype.Text        `json:"column_4"`
	Column5 pgtype.Text        `json:"column_5"`
	Column6 pgtype.Text        `json:"column_6"`
	Column7 pgtype.Timestamptz `json:"column_7"`
}

func (q *Queries) InsertLegalDocument(ctx context.Context, arg InsertLegalDocumentParams) error {
	_, err := q.db.Exec(ctx, insertLegalDocument,
		arg.Column1,
		arg.Column2,
		arg.Column3,
		arg.Column4,
		arg.Column5,
		arg.Column6,
		arg.Column7,
	)
	return err
}

const updateLegalDocument = `-- name: UpdateLegalDocument :execrows
UPDATE
    legal_documents
SET
    document_name = $1,
    file_name = $2,
    document_type = $3,
    description = $4,
    updated_at = $5
WHERE
    id = $6
AND
    deleted_at IS NULL
`

type UpdateLegalDocumentParams struct {
	DocumentName pgtype.Text        `json:"document_name"`
	FileName     pgtype.Text        `json:"file_name"`
	DocumentType pgtype.Text        `json:"document_type"`
	Description  pgtype.Text        `json:"description"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
	ID           pgtype.Text        `json:"id"`
}

func (q *Queries) UpdateLegalDocument(ctx context.Context, arg UpdateLegalDocumentParams) (int64, error) {
	result, err := q.db.Exec(ctx, updateLegalDocument,
		arg.DocumentName,
		arg.FileName,
		arg.DocumentType,
		arg.Description,
		arg.UpdatedAt,
		arg.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
