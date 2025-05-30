// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: legal_document_type.sql

package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteLegalDocType = `-- name: DeleteLegalDocType :execrows
UPDATE
    legal_document_types
SET
    deleted_at = $1
WHERE
    id = $2
AND
    deleted_at IS NULL
`

type DeleteLegalDocTypeParams struct {
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
	ID        string             `json:"id"`
}

func (q *Queries) DeleteLegalDocType(ctx context.Context, arg DeleteLegalDocTypeParams) (int64, error) {
	result, err := q.db.Exec(ctx, deleteLegalDocType, arg.DeletedAt, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const findAllLegalDocTypes = `-- name: FindAllLegalDocTypes :many
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
    description ASC
`

type FindAllLegalDocTypesRow struct {
	ID          string             `json:"id"`
	Description string             `json:"description"`
	Author      string             `json:"author"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) FindAllLegalDocTypes(ctx context.Context) ([]FindAllLegalDocTypesRow, error) {
	rows, err := q.db.Query(ctx, findAllLegalDocTypes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FindAllLegalDocTypesRow{}
	for rows.Next() {
		var i FindAllLegalDocTypesRow
		if err := rows.Scan(
			&i.ID,
			&i.Description,
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

const findLegalDocTypeByID = `-- name: FindLegalDocTypeByID :one
SELECT
    id,
    description,
    author,
    created_at,
    updated_at
FROM
    legal_document_types
WHERE
    id = $1
AND
    deleted_at IS NULL
`

type FindLegalDocTypeByIDRow struct {
	ID          string             `json:"id"`
	Description string             `json:"description"`
	Author      string             `json:"author"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) FindLegalDocTypeByID(ctx context.Context, id string) (FindLegalDocTypeByIDRow, error) {
	row := q.db.QueryRow(ctx, findLegalDocTypeByID, id)
	var i FindLegalDocTypeByIDRow
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Author,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertLegalDocType = `-- name: InsertLegalDocType :exec
INSERT INTO legal_document_types(id, description, author, created_at)
VALUES (
    $1,
    $2,
    $3,
    $4
)
`

type InsertLegalDocTypeParams struct {
	ID          string             `json:"id"`
	Description string             `json:"description"`
	Author      string             `json:"author"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) InsertLegalDocType(ctx context.Context, arg InsertLegalDocTypeParams) error {
	_, err := q.db.Exec(ctx, insertLegalDocType,
		arg.ID,
		arg.Description,
		arg.Author,
		arg.CreatedAt,
	)
	return err
}

const updateLegalDocType = `-- name: UpdateLegalDocType :execrows
UPDATE
    legal_document_types
SET
    description = $1,
    updated_at = $2
WHERE
    id = $3
AND
    deleted_at IS NULL
`

type UpdateLegalDocTypeParams struct {
	Description string             `json:"description"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	ID          string             `json:"id"`
}

func (q *Queries) UpdateLegalDocType(ctx context.Context, arg UpdateLegalDocTypeParams) (int64, error) {
	result, err := q.db.Exec(ctx, updateLegalDocType, arg.Description, arg.UpdatedAt, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
