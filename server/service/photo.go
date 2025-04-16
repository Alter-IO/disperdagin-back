package service

import (
	"alter-io-go/helpers/derrors"
	helpers "alter-io-go/helpers/ulid"
	"alter-io-go/repositories/postgresql"
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Service) GetAllPhotos(ctx context.Context) ([]postgresql.FindAllPhotosRow, error) {
	photos, err := s.repo.FindAllPhotos(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return photos, nil
}

func (s *Service) GetPhotoByID(ctx context.Context, id string) (postgresql.FindPhotoByIDRow, error) {
	photo, err := s.repo.FindPhotoByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return photo, derrors.NewErrorf(derrors.ErrorCodeNotFound, "foto tidak ditemukan")
		}
		return photo, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return photo, nil
}

func (s *Service) GetPhotosByCategory(ctx context.Context, categoryID string) ([]postgresql.FindPhotosByCategoryRow, error) {
	// Optionally, verify the category exists first
	_, err := s.repo.FindPhotoCategoryByID(ctx, categoryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, derrors.NewErrorf(derrors.ErrorCodeNotFound, "kategori foto tidak ditemukan")
		}
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	photos, err := s.repo.FindPhotosByCategory(ctx, categoryID)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return photos, nil
}

func validateCreatePhoto(data postgresql.InsertPhotoParams) error {
	if data.CategoryID == "" {
		return errors.New("id kategori wajib diisi")
	}

	if data.Title == "" {
		return errors.New("judul foto wajib diisi")
	}

	if !data.FileUrl.Valid || data.FileUrl.String == "" {
		return errors.New("file foto wajib diisi")
	}

	if data.Author == "" {
		return errors.New("penulis wajib diisi")
	}

	return nil
}

func (s *Service) CreatePhoto(ctx context.Context, data postgresql.InsertPhotoParams) error {
	if err := validateCreatePhoto(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	// Verify the category exists
	_, err := s.repo.FindPhotoCategoryByID(ctx, data.CategoryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return derrors.NewErrorf(derrors.ErrorCodeNotFound, "kategori foto tidak ditemukan")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	params := postgresql.InsertPhotoParams{
		ID:          helpers.GenerateID(),
		CategoryID:  data.CategoryID,
		Title:       data.Title,
		FileUrl:     data.FileUrl,
		Description: data.Description,
		Author:      data.Author,
		CreatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertPhoto(ctx, params); err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func validateUpdatePhoto(data postgresql.UpdatePhotoParams) error {
	if data.CategoryID == "" {
		return errors.New("id kategori wajib diisi")
	}

	if data.Title == "" {
		return errors.New("judul foto wajib diisi")
	}

	return nil
}

func (s *Service) UpdatePhoto(ctx context.Context, data postgresql.UpdatePhotoParams) error {
	if err := validateUpdatePhoto(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	// Verify the category exists
	_, err := s.repo.FindPhotoCategoryByID(ctx, data.CategoryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return derrors.NewErrorf(derrors.ErrorCodeNotFound, "kategori foto tidak ditemukan")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	params := postgresql.UpdatePhotoParams{
		ID:          data.ID,
		CategoryID:  data.CategoryID,
		Title:       data.Title,
		Description: data.Description,
		FileUrl:     data.FileUrl,
		UpdatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdatePhoto(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "foto tidak ditemukan")
	}

	return nil
}

func (s *Service) DeletePhoto(ctx context.Context, id string) error {
	params := postgresql.DeletePhotoParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeletePhoto(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "foto tidak ditemukan")
	}

	return nil
}
