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

func (s *Service) GetAllCommodityTypes(ctx context.Context) ([]postgresql.FindAllCommodityTypesRow, error) {
	commodityTypes, err := s.repo.FindAllCommodityTypes(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return commodityTypes, nil
}

func (s *Service) GetCommodityTypeByID(ctx context.Context, id string) (postgresql.FindCommodityTypeByIDRow, error) {
	commodityType, err := s.repo.FindCommodityTypeByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return commodityType, derrors.NewErrorf(derrors.ErrorCodeNotFound, "tipe komoditas tidak ditemukan")
		}
		return commodityType, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return commodityType, nil
}

func (s *Service) CreateCommodityType(ctx context.Context, data postgresql.InsertCommodityTypeParams) error {
	if data.Description == "" {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "deskripsi wajib diisi")
	}

	if data.Author == "" {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "author wajib diisi")
	}

	params := postgresql.InsertCommodityTypeParams{
		ID:          helpers.GenerateID(),
		Description: data.Description,
		Author:      data.Author,
		CreatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertCommodityType(ctx, params); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return derrors.NewErrorf(derrors.ErrorCodeDuplicate, "tipe komoditas sudah ada")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func (s *Service) UpdateCommodityType(ctx context.Context, params postgresql.UpdateCommodityTypeParams) error {
	if params.Description == "" {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "deskripsi wajib diisi")
	}

	// Set the updated time
	params.UpdatedAt = pgtype.Timestamptz{Time: time.Now(), Valid: true}

	rowsAffected, err := s.repo.UpdateCommodityType(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return derrors.NewErrorf(derrors.ErrorCodeDuplicate, "tipe komoditas sudah ada")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "tipe komoditas tidak ditemukan")
	}

	return nil
}

func (s *Service) DeleteCommodityType(ctx context.Context, id string) error {
	params := postgresql.DeleteCommodityTypeParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeleteCommodityType(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "tipe komoditas tidak ditemukan")
	}

	return nil
}
