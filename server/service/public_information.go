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

func (s *Service) GetAllPublicInfo(ctx context.Context) ([]postgresql.FindAllPublicInfoRow, error) {
	info, err := s.repo.FindAllPublicInfo(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return info, nil
}

func (s *Service) GetPublicInfoByID(ctx context.Context, id string) (postgresql.FindPublicInfoByIDRow, error) {
	info, err := s.repo.FindPublicInfoByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return info, derrors.NewErrorf(derrors.ErrorCodeNotFound, "informasi publik tidak ditemukan")
		}
		return info, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return info, nil
}

func (s *Service) GetPublicInfoByType(ctx context.Context, infoType string) ([]postgresql.FindPublicInfoByTypeRow, error) {
	info, err := s.repo.FindPublicInfoByType(ctx, infoType)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return info, nil
}

func validateCreatePublicInfo(data postgresql.InsertPublicInfoParams) error {
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

func (s *Service) CreatePublicInfo(ctx context.Context, data postgresql.InsertPublicInfoParams) error {
	if err := validateCreatePublicInfo(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	// Optionally, we could check if the public info type exists
	// This would improve data integrity but requires additional query

	params := postgresql.InsertPublicInfoParams{
		ID:             helpers.GenerateID(),
		DocumentName:   data.DocumentName,
		FileName:       data.FileName,
		PublicInfoType: data.PublicInfoType,
		Description:    data.Description,
		Author:         data.Author,
		CreatedAt:      pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertPublicInfo(ctx, params); err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func validateUpdatePublicInfo(data postgresql.UpdatePublicInfoParams) error {
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

func (s *Service) UpdatePublicInfo(ctx context.Context, data postgresql.UpdatePublicInfoParams) error {
	if err := validateUpdatePublicInfo(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	// Optionally, we could check if the public info type exists
	// This would improve data integrity but requires additional query

	params := postgresql.UpdatePublicInfoParams{
		ID:             data.ID,
		DocumentName:   data.DocumentName,
		FileName:       data.FileName,
		PublicInfoType: data.PublicInfoType,
		Description:    data.Description,
		UpdatedAt:      pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdatePublicInfo(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "informasi publik tidak ditemukan")
	}

	return nil
}

func (s *Service) DeletePublicInfo(ctx context.Context, id string) error {
	params := postgresql.DeletePublicInfoParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeletePublicInfo(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "informasi publik tidak ditemukan")
	}

	return nil
}
