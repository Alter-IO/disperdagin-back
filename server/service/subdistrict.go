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

func (s *Service) GetAllSubdistricts(ctx context.Context) ([]postgresql.FindAllSubdistrictsRow, error) {
	subdistricts, err := s.repo.FindAllSubdistricts(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return subdistricts, nil
}

func (s *Service) GetSubdistrictByID(ctx context.Context, id string) (postgresql.FindSubdistrictByIDRow, error) {
	subdistrict, err := s.repo.FindSubdistrictByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return subdistrict, derrors.NewErrorf(derrors.ErrorCodeNotFound, "kecamatan tidak ditemukan")
		}
		return subdistrict, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return subdistrict, nil
}

func validateCreateSubdistrict(data postgresql.InsertSubdistrictParams) error {
	if data.Name == "" {
		return errors.New("nama kecamatan wajib diisi")
	}

	if data.Author == "" {
		return errors.New("penulis wajib diisi")
	}

	return nil
}

func (s *Service) CreateSubdistrict(ctx context.Context, data postgresql.InsertSubdistrictParams) error {
	if err := validateCreateSubdistrict(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	params := postgresql.InsertSubdistrictParams{
		ID:        helpers.GenerateID(),
		Name:      data.Name,
		Author:    data.Author,
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertSubdistrict(ctx, params); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return derrors.NewErrorf(derrors.ErrorCodeDuplicate, "kecamatan dengan nama tersebut sudah ada")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func (s *Service) UpdateSubdistrict(ctx context.Context, data postgresql.UpdateSubdistrictParams) error {
	if data.Name == "" {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "nama kecamatan wajib diisi")
	}

	params := postgresql.UpdateSubdistrictParams{
		ID:        data.ID,
		Name:      data.Name,
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdateSubdistrict(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return derrors.NewErrorf(derrors.ErrorCodeDuplicate, "kecamatan dengan nama tersebut sudah ada")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "kecamatan tidak ditemukan")
	}

	return nil
}

func (s *Service) DeleteSubdistrict(ctx context.Context, id string) error {
	// Optionally, we could check if there are any villages in this subdistrict
	// and prevent deletion if there are, to maintain data integrity

	params := postgresql.DeleteSubdistrictParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeleteSubdistrict(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "kecamatan tidak ditemukan")
	}

	return nil
}
