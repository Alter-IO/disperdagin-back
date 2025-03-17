// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: photos.sql

package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deletePhoto = `-- name: DeletePhoto :execrows
UPDATE
    photos
SET
    deleted_at = $1
WHERE
    id = $2
AND
    deleted_at IS NULL
`

type DeletePhotoParams struct {
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
	ID        string             `json:"id"`
}

func (q *Queries) DeletePhoto(ctx context.Context, arg DeletePhotoParams) (int64, error) {
	result, err := q.db.Exec(ctx, deletePhoto, arg.DeletedAt, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const findAllPhotos = `-- name: FindAllPhotos :many
SELECT
    id,
    category_id,
    title,
    file,
    author,
    created_at
FROM
    photos
WHERE
    deleted_at IS NULL
ORDER BY
    created_at DESC
`

type FindAllPhotosRow struct {
	ID         string             `json:"id"`
	CategoryID string             `json:"category_id"`
	Title      string             `json:"title"`
	File       pgtype.Text        `json:"file"`
	Author     string             `json:"author"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) FindAllPhotos(ctx context.Context) ([]FindAllPhotosRow, error) {
	rows, err := q.db.Query(ctx, findAllPhotos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FindAllPhotosRow{}
	for rows.Next() {
		var i FindAllPhotosRow
		if err := rows.Scan(
			&i.ID,
			&i.CategoryID,
			&i.Title,
			&i.File,
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

const findPhotoByID = `-- name: FindPhotoByID :one
SELECT
    id,
    category_id,
    title,
    file,
    description,
    author,
    created_at,
    updated_at
FROM
    photos
WHERE
    id = $1
AND
    deleted_at IS NULL
`

type FindPhotoByIDRow struct {
	ID          string             `json:"id"`
	CategoryID  string             `json:"category_id"`
	Title       string             `json:"title"`
	File        pgtype.Text        `json:"file"`
	Description pgtype.Text        `json:"description"`
	Author      string             `json:"author"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) FindPhotoByID(ctx context.Context, id string) (FindPhotoByIDRow, error) {
	row := q.db.QueryRow(ctx, findPhotoByID, id)
	var i FindPhotoByIDRow
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Title,
		&i.File,
		&i.Description,
		&i.Author,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findPhotosByCategory = `-- name: FindPhotosByCategory :many
SELECT
    id,
    category_id,
    title,
    file,
    author,
    created_at
FROM
    photos
WHERE
    category_id = $1
AND
    deleted_at IS NULL
ORDER BY
    created_at DESC
`

type FindPhotosByCategoryRow struct {
	ID         string             `json:"id"`
	CategoryID string             `json:"category_id"`
	Title      string             `json:"title"`
	File       pgtype.Text        `json:"file"`
	Author     string             `json:"author"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) FindPhotosByCategory(ctx context.Context, categoryID string) ([]FindPhotosByCategoryRow, error) {
	rows, err := q.db.Query(ctx, findPhotosByCategory, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FindPhotosByCategoryRow{}
	for rows.Next() {
		var i FindPhotosByCategoryRow
		if err := rows.Scan(
			&i.ID,
			&i.CategoryID,
			&i.Title,
			&i.File,
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

const insertPhoto = `-- name: InsertPhoto :exec
INSERT INTO photos(id, category_id, title, file, description, author, created_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
`

type InsertPhotoParams struct {
	ID          string             `json:"id"`
	CategoryID  string             `json:"category_id"`
	Title       string             `json:"title"`
	File        pgtype.Text        `json:"file"`
	Description pgtype.Text        `json:"description"`
	Author      string             `json:"author"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) InsertPhoto(ctx context.Context, arg InsertPhotoParams) error {
	_, err := q.db.Exec(ctx, insertPhoto,
		arg.ID,
		arg.CategoryID,
		arg.Title,
		arg.File,
		arg.Description,
		arg.Author,
		arg.CreatedAt,
	)
	return err
}

const updatePhoto = `-- name: UpdatePhoto :execrows
UPDATE
    photos
SET
    category_id = $1,
    title = $2,
    description = $3,
    file = $4,
    updated_at = $5
WHERE
    id = $6
AND
    deleted_at IS NULL
`

type UpdatePhotoParams struct {
	CategoryID  string             `json:"category_id"`
	Title       string             `json:"title"`
	Description pgtype.Text        `json:"description"`
	File        pgtype.Text        `json:"file"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	ID          string             `json:"id"`
}

func (q *Queries) UpdatePhoto(ctx context.Context, arg UpdatePhotoParams) (int64, error) {
	result, err := q.db.Exec(ctx, updatePhoto,
		arg.CategoryID,
		arg.Title,
		arg.Description,
		arg.File,
		arg.UpdatedAt,
		arg.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
