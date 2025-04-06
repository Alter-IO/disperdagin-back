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

func (s *Service) GetSectors(ctx context.Context) ([]postgresql.FindAllSectorsRow, error) {
	return s.repo.FindAllSectors(ctx)
}

func (s *Service) GetSector(ctx context.Context, id string) (postgresql.FindSectorByIDRow, error) {
	sector, err := s.repo.FindSectorByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return sector, derrors.NewErrorf(derrors.ErrorCodeNotFound, "sektor tidak ditemukan")
		}
		return sector, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return sector, nil
}

func validateCreateSector(data postgresql.InsertSectorParams) error {
	if data.Name == "" {
		return errors.New("nama wajib diisi")
	}

	if data.Description.String == "" {
		return errors.New("deskripsi wajib diisi")
	}

	return nil
}

func (s *Service) CreateSector(ctx context.Context, data postgresql.InsertSectorParams) error {
	if err := validateCreateSector(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	params := postgresql.InsertSectorParams{
		ID:          helpers.GenerateID(),
		Name:        data.Name,
		Description: data.Description,
		Author:      data.Author,
		CreatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
	if err := s.repo.InsertSector(ctx, params); err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func validateUpdateSector(data postgresql.UpdateSectorParams) error {
	if data.Name == "" {
		return errors.New("nama wajib diisi")
	}

	if data.Description.String == "" {
		return errors.New("deskripsi wajib diisi")
	}

	return nil
}

func (s *Service) UpdateSector(ctx context.Context, id string, data postgresql.UpdateSectorParams) error {
	if err := validateUpdateSector(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	params := postgresql.UpdateSectorParams{
		ID:          id,
		Name:        data.Name,
		Description: data.Description,
		Author:      data.Author,
		UpdatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdateSector(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "sektor tidak dapat ditemukan")
	}

	return nil
}

func (s *Service) DeleteSector(ctx context.Context, id string) error {
	params := postgresql.DeleteSectorParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	_, err := s.repo.DeleteSector(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}
