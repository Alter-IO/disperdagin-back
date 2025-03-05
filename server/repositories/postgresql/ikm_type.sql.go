// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: ikm_type.sql

package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteIKMType = `-- name: DeleteIKMType :execrows
UPDATE
    ikm_types
SET
    deleted_at = $1
WHERE
    id = $2
AND
    deleted_at IS NULL
`

type DeleteIKMTypeParams struct {
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
	ID        pgtype.Text        `json:"id"`
}

func (q *Queries) DeleteIKMType(ctx context.Context, arg DeleteIKMTypeParams) (int64, error) {
	result, err := q.db.Exec(ctx, deleteIKMType, arg.DeletedAt, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const findAllIKMTypes = `-- name: FindAllIKMTypes :many
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
    document_name ASC
`

type FindAllIKMTypesRow struct {
	ID             string             `json:"id"`
	DocumentName   string             `json:"document_name"`
	FileName       string             `json:"file_name"`
	PublicInfoType string             `json:"public_info_type"`
	Author         string             `json:"author"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) FindAllIKMTypes(ctx context.Context) ([]FindAllIKMTypesRow, error) {
	rows, err := q.db.Query(ctx, findAllIKMTypes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FindAllIKMTypesRow{}
	for rows.Next() {
		var i FindAllIKMTypesRow
		if err := rows.Scan(
			&i.ID,
			&i.DocumentName,
			&i.FileName,
			&i.PublicInfoType,
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

const findIKMTypeByID = `-- name: FindIKMTypeByID :one
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
    deleted_at IS NULL
`

type FindIKMTypeByIDRow struct {
	ID             string             `json:"id"`
	DocumentName   string             `json:"document_name"`
	FileName       string             `json:"file_name"`
	PublicInfoType string             `json:"public_info_type"`
	Description    pgtype.Text        `json:"description"`
	Author         string             `json:"author"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
	UpdatedAt      pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) FindIKMTypeByID(ctx context.Context, dollar_1 pgtype.Text) (FindIKMTypeByIDRow, error) {
	row := q.db.QueryRow(ctx, findIKMTypeByID, dollar_1)
	var i FindIKMTypeByIDRow
	err := row.Scan(
		&i.ID,
		&i.DocumentName,
		&i.FileName,
		&i.PublicInfoType,
		&i.Description,
		&i.Author,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findIKMTypesByInfoType = `-- name: FindIKMTypesByInfoType :many
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
    document_name ASC
`

type FindIKMTypesByInfoTypeRow struct {
	ID             string             `json:"id"`
	DocumentName   string             `json:"document_name"`
	FileName       string             `json:"file_name"`
	PublicInfoType string             `json:"public_info_type"`
	Author         string             `json:"author"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) FindIKMTypesByInfoType(ctx context.Context, dollar_1 pgtype.Text) ([]FindIKMTypesByInfoTypeRow, error) {
	rows, err := q.db.Query(ctx, findIKMTypesByInfoType, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FindIKMTypesByInfoTypeRow{}
	for rows.Next() {
		var i FindIKMTypesByInfoTypeRow
		if err := rows.Scan(
			&i.ID,
			&i.DocumentName,
			&i.FileName,
			&i.PublicInfoType,
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

const insertIKMType = `-- name: InsertIKMType :exec
INSERT INTO ikm_types(id, document_name, file_name, public_info_type, description, author, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type InsertIKMTypeParams struct {
	Column1 pgtype.Text        `json:"column_1"`
	Column2 pgtype.Text        `json:"column_2"`
	Column3 pgtype.Text        `json:"column_3"`
	Column4 pgtype.Text        `json:"column_4"`
	Column5 pgtype.Text        `json:"column_5"`
	Column6 pgtype.Text        `json:"column_6"`
	Column7 pgtype.Timestamptz `json:"column_7"`
}

func (q *Queries) InsertIKMType(ctx context.Context, arg InsertIKMTypeParams) error {
	_, err := q.db.Exec(ctx, insertIKMType,
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

const updateIKMType = `-- name: UpdateIKMType :execrows
UPDATE
    ikm_types
SET
    document_name = $1,
    file_name = $2,
    public_info_type = $3,
    description = $4,
    updated_at = $5
WHERE
    id = $6
AND
    deleted_at IS NULL
`

type UpdateIKMTypeParams struct {
	DocumentName   pgtype.Text        `json:"document_name"`
	FileName       pgtype.Text        `json:"file_name"`
	PublicInfoType pgtype.Text        `json:"public_info_type"`
	Description    pgtype.Text        `json:"description"`
	UpdatedAt      pgtype.Timestamptz `json:"updated_at"`
	ID             pgtype.Text        `json:"id"`
}

func (q *Queries) UpdateIKMType(ctx context.Context, arg UpdateIKMTypeParams) (int64, error) {
	result, err := q.db.Exec(ctx, updateIKMType,
		arg.DocumentName,
		arg.FileName,
		arg.PublicInfoType,
		arg.Description,
		arg.UpdatedAt,
		arg.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
