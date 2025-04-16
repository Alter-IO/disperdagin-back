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

func (s *Service) GetAllLegalDocuments(ctx context.Context) ([]postgresql.FindAllLegalDocumentsRow, error) {
	documents, err := s.repo.FindAllLegalDocuments(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return documents, nil
}

func (s *Service) GetLegalDocumentByID(ctx context.Context, id string) (postgresql.FindLegalDocumentByIDRow, error) {
	document, err := s.repo.FindLegalDocumentByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return document, derrors.NewErrorf(derrors.ErrorCodeNotFound, "dokumen hukum tidak ditemukan")
		}
		return document, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return document, nil
}

func (s *Service) GetLegalDocumentsByType(ctx context.Context, docType string) ([]postgresql.FindLegalDocumentsByTypeRow, error) {
	documents, err := s.repo.FindLegalDocumentsByType(ctx, docType)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return documents, nil
}

func validateCreateLegalDocument(data postgresql.InsertLegalDocumentParams) error {
	if data.DocumentName == "" {
		return errors.New("nama dokumen wajib diisi")
	}

	if data.FileUrl == "" {
		return errors.New("nama file wajib diisi")
	}

	if data.DocumentType == "" {
		return errors.New("tipe dokumen wajib diisi")
	}

	if data.Author == "" {
		return errors.New("penulis wajib diisi")
	}

	return nil
}

func (s *Service) CreateLegalDocument(ctx context.Context, data postgresql.InsertLegalDocumentParams) error {
	if err := validateCreateLegalDocument(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	// Optionally: Check if the document type exists
	// This would improve data integrity but requires additional query

	params := postgresql.InsertLegalDocumentParams{
		ID:           helpers.GenerateID(),
		DocumentName: data.DocumentName,
		FileUrl:      data.FileUrl,
		DocumentType: data.DocumentType,
		Description:  data.Description,
		Author:       data.Author,
		CreatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertLegalDocument(ctx, params); err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func validateUpdateLegalDocument(data postgresql.UpdateLegalDocumentParams) error {
	if data.DocumentName == "" {
		return errors.New("nama dokumen wajib diisi")
	}

	if data.FileUrl == "" {
		return errors.New("nama file wajib diisi")
	}

	if data.DocumentType == "" {
		return errors.New("tipe dokumen wajib diisi")
	}

	return nil
}

func (s *Service) UpdateLegalDocument(ctx context.Context, data postgresql.UpdateLegalDocumentParams) error {
	if err := validateUpdateLegalDocument(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	// Optionally: Check if the document type exists
	// This would improve data integrity but requires additional query

	params := postgresql.UpdateLegalDocumentParams{
		ID:           data.ID,
		DocumentName: data.DocumentName,
		FileUrl:      data.FileUrl,
		DocumentType: data.DocumentType,
		Description:  data.Description,
		UpdatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdateLegalDocument(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "dokumen hukum tidak ditemukan")
	}

	return nil
}

func (s *Service) DeleteLegalDocument(ctx context.Context, id string) error {
	params := postgresql.DeleteLegalDocumentParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeleteLegalDocument(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "dokumen hukum tidak ditemukan")
	}

	return nil
}
