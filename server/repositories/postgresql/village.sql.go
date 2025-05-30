// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: village.sql

package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteVillage = `-- name: DeleteVillage :execrows
UPDATE
    villages
SET
    deleted_at = $1
WHERE
    id = $2
AND
    deleted_at IS NULL
`

type DeleteVillageParams struct {
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
	ID        string             `json:"id"`
}

func (q *Queries) DeleteVillage(ctx context.Context, arg DeleteVillageParams) (int64, error) {
	result, err := q.db.Exec(ctx, deleteVillage, arg.DeletedAt, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const findAllVillages = `-- name: FindAllVillages :many
SELECT
    id,
    subdistrict_id,
    name,
    author,
    created_at
FROM
    villages
WHERE
    deleted_at IS NULL
ORDER BY
    name ASC
`

type FindAllVillagesRow struct {
	ID            string             `json:"id"`
	SubdistrictID string             `json:"subdistrict_id"`
	Name          string             `json:"name"`
	Author        string             `json:"author"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) FindAllVillages(ctx context.Context) ([]FindAllVillagesRow, error) {
	rows, err := q.db.Query(ctx, findAllVillages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FindAllVillagesRow{}
	for rows.Next() {
		var i FindAllVillagesRow
		if err := rows.Scan(
			&i.ID,
			&i.SubdistrictID,
			&i.Name,
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

const findVillageByID = `-- name: FindVillageByID :one
SELECT
    id,
    subdistrict_id,
    name,
    author,
    created_at,
    updated_at
FROM
    villages
WHERE
    id = $1
AND
    deleted_at IS NULL
`

type FindVillageByIDRow struct {
	ID            string             `json:"id"`
	SubdistrictID string             `json:"subdistrict_id"`
	Name          string             `json:"name"`
	Author        string             `json:"author"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) FindVillageByID(ctx context.Context, id string) (FindVillageByIDRow, error) {
	row := q.db.QueryRow(ctx, findVillageByID, id)
	var i FindVillageByIDRow
	err := row.Scan(
		&i.ID,
		&i.SubdistrictID,
		&i.Name,
		&i.Author,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findVillagesBySubdistrict = `-- name: FindVillagesBySubdistrict :many
SELECT
    id,
    subdistrict_id,
    name,
    author,
    created_at
FROM
    villages
WHERE
    subdistrict_id = $1
AND
    deleted_at IS NULL
ORDER BY
    name ASC
`

type FindVillagesBySubdistrictRow struct {
	ID            string             `json:"id"`
	SubdistrictID string             `json:"subdistrict_id"`
	Name          string             `json:"name"`
	Author        string             `json:"author"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) FindVillagesBySubdistrict(ctx context.Context, subdistrictID string) ([]FindVillagesBySubdistrictRow, error) {
	rows, err := q.db.Query(ctx, findVillagesBySubdistrict, subdistrictID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FindVillagesBySubdistrictRow{}
	for rows.Next() {
		var i FindVillagesBySubdistrictRow
		if err := rows.Scan(
			&i.ID,
			&i.SubdistrictID,
			&i.Name,
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

const insertVillage = `-- name: InsertVillage :exec
INSERT INTO villages(id, subdistrict_id, name, author, created_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
`

type InsertVillageParams struct {
	ID            string             `json:"id"`
	SubdistrictID string             `json:"subdistrict_id"`
	Name          string             `json:"name"`
	Author        string             `json:"author"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) InsertVillage(ctx context.Context, arg InsertVillageParams) error {
	_, err := q.db.Exec(ctx, insertVillage,
		arg.ID,
		arg.SubdistrictID,
		arg.Name,
		arg.Author,
		arg.CreatedAt,
	)
	return err
}

const updateVillage = `-- name: UpdateVillage :execrows
UPDATE
    villages
SET
    subdistrict_id = $1,
    name = $2,
    updated_at = $3
WHERE
    id = $4
AND
    deleted_at IS NULL
`

type UpdateVillageParams struct {
	SubdistrictID string             `json:"subdistrict_id"`
	Name          string             `json:"name"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
	ID            string             `json:"id"`
}

func (q *Queries) UpdateVillage(ctx context.Context, arg UpdateVillageParams) (int64, error) {
	result, err := q.db.Exec(ctx, updateVillage,
		arg.SubdistrictID,
		arg.Name,
		arg.UpdatedAt,
		arg.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
