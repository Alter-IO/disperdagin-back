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

func (s *Service) GetAllIKMs(ctx context.Context) ([]postgresql.FindAllIKMsRow, error) {
	ikms, err := s.repo.FindAllIKMs(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return ikms, nil
}

func (s *Service) GetIKMByID(ctx context.Context, id string) (postgresql.IKMRow, error) {
	ikm, err := s.repo.IKM(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ikm, derrors.NewErrorf(derrors.ErrorCodeNotFound, "data IKM tidak ditemukan")
		}
		return ikm, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return ikm, nil
}

func (s *Service) GetIKMsByVillage(ctx context.Context, villageID string) ([]postgresql.FindIKMsByVillageRow, error) {
	ikms, err := s.repo.FindIKMsByVillage(ctx, villageID)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return ikms, nil
}

func (s *Service) GetIKMsByBusinessType(ctx context.Context, businessType string) ([]postgresql.FindIKMsByBusinessTypeRow, error) {
	ikms, err := s.repo.FindIKMsByBusinessType(ctx, businessType)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return ikms, nil
}

func validateCreateIKM(data postgresql.InsertIKMParams) error {
	if data.Description == "" {
		return errors.New("deskripsi IKM wajib diisi")
	}

	if data.VillageID == "" {
		return errors.New("ID desa wajib diisi")
	}

	if data.BusinessType == "" {
		return errors.New("tipe bisnis wajib diisi")
	}

	if data.Author == "" {
		return errors.New("author wajib diisi")
	}

	return nil
}

func (s *Service) CreateIKM(ctx context.Context, data postgresql.InsertIKMParams) error {
	if err := validateCreateIKM(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	// Verify if the village exists
	_, err := s.repo.FindVillageByID(ctx, data.VillageID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return derrors.NewErrorf(derrors.ErrorCodeNotFound, "desa tidak ditemukan")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	params := postgresql.InsertIKMParams{
		ID:           helpers.GenerateID(),
		Description:  data.Description,
		VillageID:    data.VillageID,
		BusinessType: data.BusinessType,
		Author:       data.Author,
		CreatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertIKM(ctx, params); err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func validateUpdateIKM(data postgresql.UpdateIKMParams) error {
	if data.Description == "" {
		return errors.New("deskripsi IKM wajib diisi")
	}

	if data.VillageID == "" {
		return errors.New("ID desa wajib diisi")
	}

	if data.BusinessType == "" {
		return errors.New("tipe bisnis wajib diisi")
	}

	return nil
}

func (s *Service) UpdateIKM(ctx context.Context, data postgresql.UpdateIKMParams) error {
	if err := validateUpdateIKM(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	// Verify if the village exists
	_, err := s.repo.FindVillageByID(ctx, data.VillageID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return derrors.NewErrorf(derrors.ErrorCodeNotFound, "desa tidak ditemukan")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	params := postgresql.UpdateIKMParams{
		ID:           data.ID,
		Description:  data.Description,
		VillageID:    data.VillageID,
		BusinessType: data.BusinessType,
		UpdatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdateIKM(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "data IKM tidak ditemukan")
	}

	return nil
}

func (s *Service) DeleteIKM(ctx context.Context, id string) error {
	params := postgresql.DeleteIKMParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeleteIKM(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "data IKM tidak ditemukan")
	}

	return nil
}
