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

func (s *Service) GetUsers(ctx context.Context) ([]postgresql.FindUsersRow, error) {
	users, err := s.repo.FindUsers(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return users, nil
}

func (s *Service) GetUserByID(ctx context.Context, id string) (postgresql.FindUserByIDRow, error) {
	user, err := s.repo.FindUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, derrors.NewErrorf(derrors.ErrorCodeNotFound, "uraian jenis keigatan tidak ditemukan")
		}
		return user, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return user, nil
}

func validateCreateUser(data postgresql.InsertUserParams) error {
	if data.Username == "" {
		return errors.New("username wajib di isi")
	}

	if data.Password == "" {
		return errors.New("password wajib di isi")
	}

	return nil
}

func (s *Service) CreateUser(ctx context.Context, data postgresql.InsertUserParams) error {
	if err := validateCreateUser(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	params := postgresql.InsertUserParams{
		ID:        helpers.GenerateID(),
		RoleID:    data.RoleID,
		Username:  data.Username,
		Password:  data.Password,
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
	if err := s.repo.InsertUser(ctx, params); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return derrors.NewErrorf(derrors.ErrorCodeDuplicate, "username sudah ada")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	params := postgresql.DeleteUserParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
	rowAffected, err := s.repo.DeleteUser(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "user tidak ditemukan")
	}

	return nil
}
