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

func (s *Service) GetAllVisionMissions(ctx context.Context) ([]postgresql.FindAllVisionMissionsRow, error) {
	visionMissions, err := s.repo.FindAllVisionMissions(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return visionMissions, nil
}

func (s *Service) GetLatestVisionMission(ctx context.Context) (postgresql.FindLatestVisionMissionRow, error) {
	visionMission, err := s.repo.FindLatestVisionMission(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return visionMission, derrors.NewErrorf(derrors.ErrorCodeNotFound, "visi dan misi tidak ditemukan")
		}
		return visionMission, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return visionMission, nil
}

func (s *Service) GetVisionMissionByID(ctx context.Context, id string) (postgresql.FindVisionMissionByIDRow, error) {
	visionMission, err := s.repo.FindVisionMissionByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return visionMission, derrors.NewErrorf(derrors.ErrorCodeNotFound, "visi dan misi tidak ditemukan")
		}
		return visionMission, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return visionMission, nil
}

func validateVisionMission(data postgresql.InsertVisionMissionParams) error {
	if data.Vision == "" {
		return errors.New("visi wajib diisi")
	}

	if data.Mission == "" {
		return errors.New("misi wajib diisi")
	}

	if data.Author == "" {
		return errors.New("penulis wajib diisi")
	}

	return nil
}

func (s *Service) CreateVisionMission(ctx context.Context, data postgresql.InsertVisionMissionParams) error {
	if err := validateVisionMission(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	params := postgresql.InsertVisionMissionParams{
		ID:        helpers.GenerateID(),
		Vision:    data.Vision,
		Mission:   data.Mission,
		Author:    data.Author,
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertVisionMission(ctx, params); err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func (s *Service) UpdateVisionMission(ctx context.Context, id string, data postgresql.UpdateVisionMissionParams) error {
	if data.Vision == "" {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "visi wajib diisi")
	}

	if data.Mission == "" {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "misi wajib diisi")
	}

	params := postgresql.UpdateVisionMissionParams{
		ID:        id,
		Vision:    data.Vision,
		Mission:   data.Mission,
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdateVisionMission(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "visi dan misi tidak ditemukan")
	}

	return nil
}

func (s *Service) DeleteVisionMission(ctx context.Context, id string) error {
	params := postgresql.DeleteVisionMissionParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeleteVisionMission(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "visi dan misi tidak ditemukan")
	}

	return nil
}
