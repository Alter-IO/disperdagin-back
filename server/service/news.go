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

func (s *Service) GetAllNews(ctx context.Context) ([]postgresql.FindAllNewsRow, error) {
	return s.repo.FindAllNews(ctx)
}

func (s *Service) GetNewsById(ctx context.Context, id string) (postgresql.FindNewsByIdRow, error) {
	news, err := s.repo.FindNewsById(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return news, derrors.NewErrorf(derrors.ErrorCodeNotFound, "berita tidak ditemukan")
		}
		return news, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return news, nil
}

func validateCreateNews(data postgresql.InsertNewsParams) error {
	if data.Title == "" {
		return errors.New("judul berita wajib diisi")
	}

	if data.Content == "" {
		return errors.New("kontent berita wajib diisi")
	}

	if data.Author == "" {
		return errors.New("nama penulis wajib diisi")
	}

	return nil
}

func (s *Service) CreateNews(ctx context.Context, data postgresql.InsertNewsParams) error {
	if err := validateCreateNews(data); err != nil {
		return err
	}

	params := postgresql.InsertNewsParams{
		ID:        helpers.GenerateID(),
		Title:     data.Title,
		Content:   data.Content,
		Author:    data.Author,
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
	if err := s.repo.InsertNews(ctx, params); err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func validateUpdateNews(data postgresql.UpdateNewsParams) error {
	if data.Title == "" {
		return errors.New("judul berita wajib diisi")
	}

	if data.Content == "" {
		return errors.New("kontent berita wajib diisi")
	}

	return nil
}

func (s *Service) UpdateNews(ctx context.Context, id string, data postgresql.UpdateNewsParams) error {
	if err := validateUpdateNews(data); err != nil {
		return err
	}

	params := postgresql.UpdateNewsParams{
		ID:        id,
		Title:     data.Title,
		Content:   data.Content,
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdateNews(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "berita tidak dapat ditemukan")
	}

	return nil

}

func (s *Service) DeleteNews(ctx context.Context, id string) error {
	params := postgresql.DeleteNewsParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	_, err := s.repo.DeleteNews(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}
