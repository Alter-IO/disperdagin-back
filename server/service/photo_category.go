package service

import (
	"alter-io-go/helpers/derrors"
	helpers "alter-io-go/helpers/ulid"
	"alter-io-go/repositories/postgresql"
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Service) GetAllPhotoCategories(ctx context.Context) ([]postgresql.FindAllPhotoCategoriesRow, error) {
	categories, err := s.repo.FindAllPhotoCategories(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return categories, nil
}

func (s *Service) GetPhotoCategoryByID(ctx context.Context, id string) (postgresql.FindPhotoCategoryByIDRow, error) {
	category, err := s.repo.FindPhotoCategoryByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return category, derrors.NewErrorf(derrors.ErrorCodeNotFound, "kategori foto tidak ditemukan")
		}
		return category, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return category, nil
}

func validateCreatePhotoCategory(data postgresql.InsertPhotoCategoryParams) error {
	if data.Category == "" {
		return errors.New("nama kategori wajib diisi")
	}

	if data.Author == "" {
		return errors.New("penulis wajib diisi")
	}

	return nil
}

func (s *Service) CreatePhotoCategory(ctx context.Context, data postgresql.InsertPhotoCategoryParams) error {
	if err := validateCreatePhotoCategory(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	params := postgresql.InsertPhotoCategoryParams{
		ID:        helpers.GenerateID(),
		Category:  data.Category,
		Author:    data.Author,
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertPhotoCategory(ctx, params); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return derrors.NewErrorf(derrors.ErrorCodeDuplicate, "kategori foto sudah ada")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func (s *Service) UpdatePhotoCategory(ctx context.Context, data postgresql.UpdatePhotoCategoryParams) error {
	if data.Category == "" {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "nama kategori wajib diisi")
	}

	params := postgresql.UpdatePhotoCategoryParams{
		ID:        data.ID,
		Category:  data.Category,
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdatePhotoCategory(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return derrors.NewErrorf(derrors.ErrorCodeDuplicate, "kategori foto sudah ada")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "kategori foto tidak ditemukan")
	}

	return nil
}

func (s *Service) DeletePhotoCategory(ctx context.Context, id string) error {
	// Optionally, we could check if there are any photos in this category
	// and prevent deletion if there are, to maintain data integrity

	params := postgresql.DeletePhotoCategoryParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeletePhotoCategory(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "kategori foto tidak ditemukan")
	}

	return nil
}
