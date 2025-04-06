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

func (s *Service) GetAllVillages(ctx context.Context) ([]postgresql.FindAllVillagesRow, error) {
	villages, err := s.repo.FindAllVillages(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return villages, nil
}

func (s *Service) GetVillageByID(ctx context.Context, id string) (postgresql.FindVillageByIDRow, error) {
	village, err := s.repo.FindVillageByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return village, derrors.NewErrorf(derrors.ErrorCodeNotFound, "desa tidak ditemukan")
		}
		return village, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return village, nil
}

func (s *Service) GetVillagesBySubdistrict(ctx context.Context, subdistrictID string) ([]postgresql.FindVillagesBySubdistrictRow, error) {
	// Optionally, verify if the subdistrict exists
	_, err := s.repo.FindSubdistrictByID(ctx, subdistrictID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, derrors.NewErrorf(derrors.ErrorCodeNotFound, "kecamatan tidak ditemukan")
		}
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	villages, err := s.repo.FindVillagesBySubdistrict(ctx, subdistrictID)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return villages, nil
}

func validateCreateVillage(data postgresql.InsertVillageParams) error {
	if data.SubdistrictID == "" {
		return errors.New("id kecamatan wajib diisi")
	}

	if data.Name == "" {
		return errors.New("nama desa wajib diisi")
	}

	if data.Author == "" {
		return errors.New("penulis wajib diisi")
	}

	return nil
}

func (s *Service) CreateVillage(ctx context.Context, data postgresql.InsertVillageParams) error {
	if err := validateCreateVillage(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	// Verify if the subdistrict exists
	_, err := s.repo.FindSubdistrictByID(ctx, data.SubdistrictID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return derrors.NewErrorf(derrors.ErrorCodeNotFound, "kecamatan tidak ditemukan")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	params := postgresql.InsertVillageParams{
		ID:            helpers.GenerateID(),
		SubdistrictID: data.SubdistrictID,
		Name:          data.Name,
		Author:        data.Author,
		CreatedAt:     pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertVillage(ctx, params); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return derrors.NewErrorf(derrors.ErrorCodeDuplicate, "desa dengan nama tersebut sudah ada dalam kecamatan ini")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func validateUpdateVillage(data postgresql.UpdateVillageParams) error {
	if data.SubdistrictID == "" {
		return errors.New("id kecamatan wajib diisi")
	}

	if data.Name == "" {
		return errors.New("nama desa wajib diisi")
	}

	return nil
}

func (s *Service) UpdateVillage(ctx context.Context, data postgresql.UpdateVillageParams) error {
	if err := validateUpdateVillage(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	// Verify if the subdistrict exists
	_, err := s.repo.FindSubdistrictByID(ctx, data.SubdistrictID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return derrors.NewErrorf(derrors.ErrorCodeNotFound, "kecamatan tidak ditemukan")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	params := postgresql.UpdateVillageParams{
		ID:            data.ID,
		SubdistrictID: data.SubdistrictID,
		Name:          data.Name,
		UpdatedAt:     pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdateVillage(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return derrors.NewErrorf(derrors.ErrorCodeDuplicate, "desa dengan nama tersebut sudah ada dalam kecamatan ini")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "desa tidak ditemukan")
	}

	return nil
}

func (s *Service) DeleteVillage(ctx context.Context, id string) error {
	// Optionally, check if there are any data (like IKMs) that reference this village
	// and prevent deletion if there are, to maintain data integrity

	params := postgresql.DeleteVillageParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeleteVillage(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "desa tidak ditemukan")
	}

	return nil
}
