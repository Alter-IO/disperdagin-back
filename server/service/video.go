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

func (s *Service) GetAllVideos(ctx context.Context) ([]postgresql.FindAllVideosRow, error) {
	videos, err := s.repo.FindAllVideos(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return videos, nil
}

func (s *Service) GetVideoByID(ctx context.Context, id string) (postgresql.FindVideoByIDRow, error) {
	video, err := s.repo.FindVideoByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return video, derrors.NewErrorf(derrors.ErrorCodeNotFound, "video tidak ditemukan")
		}
		return video, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return video, nil
}

func validateCreateVideo(data postgresql.InsertVideoParams) error {
	if data.Title == "" {
		return errors.New("judul video wajib diisi")
	}

	if data.Link == "" {
		return errors.New("link video wajib diisi")
	}

	if data.Author == "" {
		return errors.New("penulis wajib diisi")
	}

	return nil
}

func (s *Service) CreateVideo(ctx context.Context, data postgresql.InsertVideoParams) error {
	if err := validateCreateVideo(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	params := postgresql.InsertVideoParams{
		ID:          helpers.GenerateID(),
		Title:       data.Title,
		Link:        data.Link,
		Description: data.Description,
		Author:      data.Author,
		CreatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertVideo(ctx, params); err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func validateUpdateVideo(data postgresql.UpdateVideoParams) error {
	if data.Title == "" {
		return errors.New("judul video wajib diisi")
	}

	if data.Link == "" {
		return errors.New("link video wajib diisi")
	}

	return nil
}

func (s *Service) UpdateVideo(ctx context.Context, data postgresql.UpdateVideoParams) error {
	if err := validateUpdateVideo(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	params := postgresql.UpdateVideoParams{
		ID:          data.ID,
		Title:       data.Title,
		Link:        data.Link,
		Description: data.Description,
		UpdatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdateVideo(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "video tidak ditemukan")
	}

	return nil
}

func (s *Service) DeleteVideo(ctx context.Context, id string) error {
	params := postgresql.DeleteVideoParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeleteVideo(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "video tidak ditemukan")
	}

	return nil
}
