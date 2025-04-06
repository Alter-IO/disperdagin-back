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

func (s *Service) GetAllMarkets(ctx context.Context) ([]postgresql.FindAllMarketsRow, error) {
	markets, err := s.repo.FindAllMarkets(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return markets, nil
}

func (s *Service) GetMarketByID(ctx context.Context, id string) (postgresql.FindMarketByIDRow, error) {
	market, err := s.repo.FindMarketByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return market, derrors.NewErrorf(derrors.ErrorCodeNotFound, "pasar tidak ditemukan")
		}
		return market, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return market, nil
}

func validateCreateMarket(data postgresql.InsertMarketParams) error {
	if data.Name == "" {
		return errors.New("nama pasar wajib diisi")
	}

	if data.Author == "" {
		return errors.New("penulis wajib diisi")
	}

	return nil
}

func (s *Service) CreateMarket(ctx context.Context, data postgresql.InsertMarketParams) error {
	if err := validateCreateMarket(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	params := postgresql.InsertMarketParams{
		ID:        helpers.GenerateID(),
		Name:      data.Name,
		Author:    data.Author,
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertMarket(ctx, params); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return derrors.NewErrorf(derrors.ErrorCodeDuplicate, "nama pasar sudah ada")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func (s *Service) UpdateMarket(ctx context.Context, data postgresql.UpdateMarketParams) error {
	if data.Name == "" {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "nama pasar wajib diisi")
	}

	params := postgresql.UpdateMarketParams{
		ID:        data.ID,
		Name:      data.Name,
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdateMarket(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return derrors.NewErrorf(derrors.ErrorCodeDuplicate, "nama pasar sudah ada")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "pasar tidak ditemukan")
	}

	return nil
}

func (s *Service) DeleteMarket(ctx context.Context, id string) error {
	// TODO: Optionally check if there are market fees associated with this market
	// This would prevent deleting markets that have associated fees

	params := postgresql.DeleteMarketParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeleteMarket(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "pasar tidak ditemukan")
	}

	return nil
}
