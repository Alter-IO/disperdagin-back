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

func (s *Service) GetAllLegalDocTypes(ctx context.Context) ([]postgresql.FindAllLegalDocTypesRow, error) {
	docTypes, err := s.repo.FindAllLegalDocTypes(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return docTypes, nil
}

func (s *Service) GetLegalDocTypeByID(ctx context.Context, id string) (postgresql.FindLegalDocTypeByIDRow, error) {
	docType, err := s.repo.FindLegalDocTypeByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return docType, derrors.NewErrorf(derrors.ErrorCodeNotFound, "tipe dokumen hukum tidak ditemukan")
		}
		return docType, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return docType, nil
}

func validateCreateLegalDocType(data postgresql.InsertLegalDocTypeParams) error {
	if data.Description == "" {
		return errors.New("deskripsi tipe dokumen wajib diisi")
	}

	if data.Author == "" {
		return errors.New("penulis wajib diisi")
	}

	return nil
}

func (s *Service) CreateLegalDocType(ctx context.Context, data postgresql.InsertLegalDocTypeParams) error {
	if err := validateCreateLegalDocType(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	params := postgresql.InsertLegalDocTypeParams{
		ID:          helpers.GenerateID(),
		Description: data.Description,
		Author:      data.Author,
		CreatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertLegalDocType(ctx, params); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return derrors.NewErrorf(derrors.ErrorCodeDuplicate, "tipe dokumen hukum sudah ada")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func (s *Service) UpdateLegalDocType(ctx context.Context, data postgresql.UpdateLegalDocTypeParams) error {
	if data.Description == "" {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "deskripsi tipe dokumen wajib diisi")
	}

	params := postgresql.UpdateLegalDocTypeParams{
		ID:          data.ID,
		Description: data.Description,
		UpdatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdateLegalDocType(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return derrors.NewErrorf(derrors.ErrorCodeDuplicate, "tipe dokumen hukum sudah ada")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "tipe dokumen hukum tidak ditemukan")
	}

	return nil
}

func (s *Service) DeleteLegalDocType(ctx context.Context, id string) error {
	// Check if there are any legal documents using this type
	// This is a data integrity check that prevents deleting types that are in use
	// Implementation would depend on additional query functionality

	params := postgresql.DeleteLegalDocTypeParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeleteLegalDocType(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "tipe dokumen hukum tidak ditemukan")
	}

	return nil
}
