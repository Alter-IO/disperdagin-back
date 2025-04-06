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

func (s *Service) GetAllMarketFees(ctx context.Context) ([]postgresql.FindAllMarketFeesRow, error) {
	fees, err := s.repo.FindAllMarketFees(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return fees, nil
}

func (s *Service) GetMarketFeeByID(ctx context.Context, id string) (postgresql.FindMarketFeeByIDRow, error) {
	fee, err := s.repo.FindMarketFeeByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fee, derrors.NewErrorf(derrors.ErrorCodeNotFound, "data retribusi pasar tidak ditemukan")
		}
		return fee, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return fee, nil
}

func (s *Service) GetMarketFeesByMarket(ctx context.Context, marketID string) ([]postgresql.FindMarketFeesByMarketRow, error) {
	fees, err := s.repo.FindMarketFeesByMarket(ctx, marketID)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return fees, nil
}

func (s *Service) GetMarketFeesByYear(ctx context.Context, year int32) ([]postgresql.FindMarketFeesByYearRow, error) {
	fees, err := s.repo.FindMarketFeesByYear(ctx, year)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return fees, nil
}

func (s *Service) GetMarketFeesBySemesterAndYear(ctx context.Context, params postgresql.FindMarketFeesBySemesterAndYearParams) ([]postgresql.FindMarketFeesBySemesterAndYearRow, error) {
	fees, err := s.repo.FindMarketFeesBySemesterAndYear(ctx, params)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return fees, nil
}

func validateCreateMarketFee(data postgresql.InsertMarketFeeParams) error {
	if data.MarketID == "" {
		return errors.New("ID pasar wajib diisi")
	}

	if data.NumPermanentKiosks < 0 {
		return errors.New("jumlah kios permanen tidak boleh negatif")
	}

	if data.NumNonPermanentKiosks < 0 {
		return errors.New("jumlah kios non-permanen tidak boleh negatif")
	}

	if !data.PermanentKioskRevenue.Valid {
		return errors.New("pendapatan kios permanen wajib diisi")
	}

	if !data.NonPermanentKioskRevenue.Valid {
		return errors.New("pendapatan kios non-permanen wajib diisi")
	}

	if data.CollectionStatus == "" {
		return errors.New("status koleksi wajib diisi")
	}

	if data.Semester == "" {
		return errors.New("semester wajib diisi")
	}

	if data.Year <= 0 {
		return errors.New("tahun wajib diisi dengan benar")
	}

	if data.Author == "" {
		return errors.New("penulis wajib diisi")
	}

	return nil
}

func (s *Service) CreateMarketFee(ctx context.Context, data postgresql.InsertMarketFeeParams) error {
	if err := validateCreateMarketFee(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	// Verify if the market exists
	_, err := s.repo.FindMarketByID(ctx, data.MarketID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return derrors.NewErrorf(derrors.ErrorCodeNotFound, "pasar tidak ditemukan")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	params := postgresql.InsertMarketFeeParams{
		ID:                       helpers.GenerateID(),
		MarketID:                 data.MarketID,
		NumPermanentKiosks:       data.NumPermanentKiosks,
		NumNonPermanentKiosks:    data.NumNonPermanentKiosks,
		PermanentKioskRevenue:    data.PermanentKioskRevenue,
		NonPermanentKioskRevenue: data.NonPermanentKioskRevenue,
		CollectionStatus:         data.CollectionStatus,
		Description:              data.Description,
		Semester:                 data.Semester,
		Year:                     data.Year,
		Author:                   data.Author,
		CreatedAt:                pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertMarketFee(ctx, params); err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func validateUpdateMarketFee(data postgresql.UpdateMarketFeeParams) error {
	if data.MarketID == "" {
		return errors.New("ID pasar wajib diisi")
	}

	if data.NumPermanentKiosks < 0 {
		return errors.New("jumlah kios permanen tidak boleh negatif")
	}

	if data.NumNonPermanentKiosks < 0 {
		return errors.New("jumlah kios non-permanen tidak boleh negatif")
	}

	if !data.PermanentKioskRevenue.Valid {
		return errors.New("pendapatan kios permanen wajib diisi")
	}

	if !data.NonPermanentKioskRevenue.Valid {
		return errors.New("pendapatan kios non-permanen wajib diisi")
	}

	if data.CollectionStatus == "" {
		return errors.New("status koleksi wajib diisi")
	}

	if data.Semester == "" {
		return errors.New("semester wajib diisi")
	}

	if data.Year <= 0 {
		return errors.New("tahun wajib diisi dengan benar")
	}

	return nil
}

func (s *Service) UpdateMarketFee(ctx context.Context, data postgresql.UpdateMarketFeeParams) error {
	if err := validateUpdateMarketFee(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	// Verify if the market exists
	_, err := s.repo.FindMarketByID(ctx, data.MarketID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return derrors.NewErrorf(derrors.ErrorCodeNotFound, "pasar tidak ditemukan")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	params := postgresql.UpdateMarketFeeParams{
		ID:                       data.ID,
		MarketID:                 data.MarketID,
		NumPermanentKiosks:       data.NumPermanentKiosks,
		NumNonPermanentKiosks:    data.NumNonPermanentKiosks,
		PermanentKioskRevenue:    data.PermanentKioskRevenue,
		NonPermanentKioskRevenue: data.NonPermanentKioskRevenue,
		CollectionStatus:         data.CollectionStatus,
		Description:              data.Description,
		Semester:                 data.Semester,
		Year:                     data.Year,
		UpdatedAt:                pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdateMarketFee(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "data retribusi pasar tidak ditemukan")
	}

	return nil
}

func (s *Service) DeleteMarketFee(ctx context.Context, id string) error {
	params := postgresql.DeleteMarketFeeParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeleteMarketFee(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "data retribusi pasar tidak ditemukan")
	}

	return nil
}
