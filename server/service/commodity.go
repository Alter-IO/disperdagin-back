package service

import (
	"alter-io-go/domain"
	"alter-io-go/helpers/derrors"
	helpers "alter-io-go/helpers/ulid"
	"alter-io-go/repositories/postgresql"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

// GetAllCommodities retrieves all commodities
func (s *Service) GetAllCommodities(ctx context.Context) ([]postgresql.FindAllCommoditiesRow, error) {
	commodities, err := s.repo.FindAllCommodities(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return commodities, nil
}

// GetLatestCommodities retrieves commodities ordered by publish date
func (s *Service) GetDailyCommodities(ctx context.Context) (postgresql.DailyCommodity, error) {
	commodity, err := s.repo.FindLatestCommodities(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return postgresql.DailyCommodity{}, nil
		}
		return postgresql.DailyCommodity{}, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return commodity, nil
}

func (s *Service) CreateDailyCommodity(ctx context.Context, data domain.CommodityDaily) error {
	if len(data.Commodities) == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "data komoditas tidak boleh kosong")
	}
	for _, commodity := range data.Commodities {
		if commodity.ID == "" {
			return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "ID komoditas wajib diisi")
		}
		if commodity.Price.LessThanOrEqual(decimal.Zero) {
			return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "Harga tidak boleh atau sama dengan Rp. 0")
		}
	}

	// masrhsalize Commodities to JSON
	commoditiesJSON, err := json.Marshal(data.Commodities)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, "gagal mengonversi data komoditas ke JSON")
	}

	params := postgresql.InsertDailyCommodityParams{
		ID:          helpers.GenerateID(),
		Commodities: commoditiesJSON,
		PublishDate: pgtype.Date{Time: data.PublishDate, Valid: true},
	}
	if err := s.repo.InsertDailyCommodity(ctx, params); err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

// GetCommoditiesByType retrieves commodities filtered by type
func (s *Service) GetCommoditiesByType(ctx context.Context, typeID string) ([]postgresql.FindCommoditiesByTypeRow, error) {
	commodities, err := s.repo.FindCommoditiesByType(ctx, typeID)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return commodities, nil
}

// GetCommodityByID retrieves a single commodity by ID
func (s *Service) GetCommodityByID(ctx context.Context, id string) (postgresql.FindCommodityByIDRow, error) {
	commodity, err := s.repo.FindCommodityByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return commodity, derrors.NewErrorf(derrors.ErrorCodeNotFound, "komoditas tidak ditemukan")
		}
		return commodity, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return commodity, nil
}

func validateCommodity(data postgresql.InsertCommodityParams) error {
	if data.Name == "" {
		return errors.New("nama komoditas wajib diisi")
	}

	if data.Unit == "" {
		return errors.New("satuan komoditas wajib diisi")
	}

	if data.CommodityTypeID == "" {
		return errors.New("tipe komoditas wajib diisi")
	}

	return nil
}

// CreateCommodity creates a new commodity
func (s *Service) CreateCommodity(ctx context.Context, data postgresql.InsertCommodityParams) error {
	if err := validateCommodity(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	// Generate a new ID
	params := postgresql.InsertCommodityParams{
		ID:              helpers.GenerateID(),
		Name:            data.Name,
		Unit:            data.Unit,
		Description:     data.Description,
		CommodityTypeID: data.CommodityTypeID,
		Author:          data.Author,
		CreatedAt:       pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertCommodity(ctx, params); err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

// UpdateCommodity updates an existing commodity
func (s *Service) UpdateCommodity(ctx context.Context, data postgresql.UpdateCommodityParams) error {
	// Validate data
	if data.Name == "" {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "nama komoditas wajib diisi")
	}

	if data.Unit == "" {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "satuan komoditas wajib diisi")
	}

	if data.CommodityTypeID == "" {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "tipe komoditas wajib diisi")
	}

	// Add current timestamp
	params := postgresql.UpdateCommodityParams{
		ID:              data.ID,
		Name:            data.Name,
		Unit:            data.Unit,
		Description:     data.Description,
		CommodityTypeID: data.CommodityTypeID,
		UpdatedAt:       pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdateCommodity(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "gagal memperbarui komoditas")
	}

	return nil
}

// DeleteCommodity soft deletes a commodity
func (s *Service) DeleteCommodity(ctx context.Context, id string) (int64, error) {
	params := postgresql.DeleteCommodityParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeleteCommodity(ctx, params)
	if err != nil {
		return 0, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return rowsAffected, nil
}
