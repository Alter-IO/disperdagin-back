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

func (s *Service) GetAllPublicInfoTypes(ctx context.Context) ([]postgresql.FindAllPublicInfoTypesRow, error) {
	infoTypes, err := s.repo.FindAllPublicInfoTypes(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return infoTypes, nil
}

func (s *Service) GetPublicInfoTypeByID(ctx context.Context, id string) (postgresql.FindPublicInfoTypeByIDRow, error) {
	infoType, err := s.repo.FindPublicInfoTypeByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return infoType, derrors.NewErrorf(derrors.ErrorCodeNotFound, "jenis informasi publik tidak ditemukan")
		}
		return infoType, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return infoType, nil
}

func validateCreatePublicInfoType(data postgresql.InsertPublicInfoTypeParams) error {
	if data.Description == "" {
		return errors.New("deskripsi jenis informasi wajib diisi")
	}

	if data.Author == "" {
		return errors.New("penulis wajib diisi")
	}

	return nil
}

func (s *Service) CreatePublicInfoType(ctx context.Context, data postgresql.InsertPublicInfoTypeParams) error {
	if err := validateCreatePublicInfoType(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	params := postgresql.InsertPublicInfoTypeParams{
		ID:          helpers.GenerateID(),
		Description: data.Description,
		Author:      data.Author,
		CreatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertPublicInfoType(ctx, params); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return derrors.NewErrorf(derrors.ErrorCodeDuplicate, "jenis informasi publik sudah ada")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func (s *Service) UpdatePublicInfoType(ctx context.Context, data postgresql.UpdatePublicInfoTypeParams) error {
	if data.Description == "" {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "deskripsi jenis informasi wajib diisi")
	}

	params := postgresql.UpdatePublicInfoTypeParams{
		ID:          data.ID,
		Description: data.Description,
		UpdatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdatePublicInfoType(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return derrors.NewErrorf(derrors.ErrorCodeDuplicate, "jenis informasi publik sudah ada")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "jenis informasi publik tidak ditemukan")
	}

	return nil
}

func (s *Service) DeletePublicInfoType(ctx context.Context, id string) error {
	// Optionally, we could check if there are any public information documents using this type
	// and prevent deletion if there are, to maintain data integrity

	params := postgresql.DeletePublicInfoTypeParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeletePublicInfoType(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "jenis informasi publik tidak ditemukan")
	}

	return nil
}
