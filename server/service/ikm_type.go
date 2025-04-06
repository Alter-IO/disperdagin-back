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

func (s *Service) GetAllIKMTypes(ctx context.Context) ([]postgresql.FindAllIKMTypesRow, error) {
	ikmTypes, err := s.repo.FindAllIKMTypes(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return ikmTypes, nil
}

func (s *Service) GetIKMTypeByID(ctx context.Context, id string) (postgresql.FindIKMTypeByIDRow, error) {
	ikmType, err := s.repo.FindIKMTypeByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ikmType, derrors.NewErrorf(derrors.ErrorCodeNotFound, "jenis IKM tidak ditemukan")
		}
		return ikmType, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return ikmType, nil
}

func (s *Service) GetIKMTypesByInfoType(ctx context.Context, infoType string) ([]postgresql.FindIKMTypesByInfoTypeRow, error) {
	ikmTypes, err := s.repo.FindIKMTypesByInfoType(ctx, infoType)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return ikmTypes, nil
}

func validateCreateIKMType(data postgresql.InsertIKMTypeParams) error {
	if data.DocumentName == "" {
		return errors.New("nama dokumen wajib diisi")
	}

	if data.FileName == "" {
		return errors.New("nama file wajib diisi")
	}

	if data.PublicInfoType == "" {
		return errors.New("tipe informasi publik wajib diisi")
	}

	if data.Author == "" {
		return errors.New("penulis wajib diisi")
	}

	return nil
}

func (s *Service) CreateIKMType(ctx context.Context, data postgresql.InsertIKMTypeParams) error {
	if err := validateCreateIKMType(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	params := postgresql.InsertIKMTypeParams{
		ID:             helpers.GenerateID(),
		DocumentName:   data.DocumentName,
		FileName:       data.FileName,
		PublicInfoType: data.PublicInfoType,
		Description:    data.Description,
		Author:         data.Author,
		CreatedAt:      pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertIKMType(ctx, params); err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func validateUpdateIKMType(data postgresql.UpdateIKMTypeParams) error {
	if data.DocumentName == "" {
		return errors.New("nama dokumen wajib diisi")
	}

	if data.FileName == "" {
		return errors.New("nama file wajib diisi")
	}

	if data.PublicInfoType == "" {
		return errors.New("tipe informasi publik wajib diisi")
	}

	return nil
}

func (s *Service) UpdateIKMType(ctx context.Context, data postgresql.UpdateIKMTypeParams) error {
	if err := validateUpdateIKMType(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	params := postgresql.UpdateIKMTypeParams{
		ID:             data.ID,
		DocumentName:   data.DocumentName,
		FileName:       data.FileName,
		PublicInfoType: data.PublicInfoType,
		Description:    data.Description,
		UpdatedAt:      pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdateIKMType(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "jenis IKM tidak ditemukan")
	}

	return nil
}

func (s *Service) DeleteIKMType(ctx context.Context, id string) error {
	params := postgresql.DeleteIKMTypeParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeleteIKMType(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "jenis IKM tidak ditemukan")
	}

	return nil
}
