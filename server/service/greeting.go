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

func (s *Service) GetAllGreetings(ctx context.Context) ([]postgresql.FindAllGreetingsRow, error) {
	greetings, err := s.repo.FindAllGreetings(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return greetings, nil
}

func (s *Service) GetLatestGreeting(ctx context.Context) (postgresql.FindLatestGreetingRow, error) {
	greeting, err := s.repo.FindLatestGreeting(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return greeting, derrors.NewErrorf(derrors.ErrorCodeNotFound, "sambutan tidak ditemukan")
		}
		return greeting, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return greeting, nil
}

func (s *Service) GetGreetingByID(ctx context.Context, id string) (postgresql.FindGreetingByIDRow, error) {
	greeting, err := s.repo.FindGreetingByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return greeting, derrors.NewErrorf(derrors.ErrorCodeNotFound, "sambutan tidak ditemukan")
		}
		return greeting, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return greeting, nil
}

func validateCreateGreeting(data postgresql.InsertGreetingParams) error {
	if data.Message == "" {
		return errors.New("pesan sambutan wajib diisi")
	}

	if data.Author == "" {
		return errors.New("penulis sambutan wajib diisi")
	}

	return nil
}

func (s *Service) CreateGreeting(ctx context.Context, data postgresql.InsertGreetingParams) error {
	if err := validateCreateGreeting(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	params := postgresql.InsertGreetingParams{
		ID:        helpers.GenerateID(),
		Message:   data.Message,
		Author:    data.Author,
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertGreeting(ctx, params); err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func (s *Service) UpdateGreeting(ctx context.Context, data postgresql.UpdateGreetingParams) error {
	if data.Message == "" {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "pesan sambutan wajib diisi")
	}

	params := postgresql.UpdateGreetingParams{
		ID:        data.ID,
		Message:   data.Message,
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdateGreeting(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "sambutan tidak ditemukan")
	}

	return nil
}

func (s *Service) DeleteGreeting(ctx context.Context, id string) error {
	params := postgresql.DeleteGreetingParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeleteGreeting(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "sambutan tidak ditemukan")
	}

	return nil
}
