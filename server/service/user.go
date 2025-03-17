package service

import (
	"alter-io-go/helpers/derrors"
	"alter-io-go/helpers/hash"
	helpers "alter-io-go/helpers/ulid"
	"alter-io-go/helpers/util"
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
			return user, derrors.NewErrorf(derrors.ErrorCodeNotFound, "user tidak ditemukan")
		}
		return user, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return user, nil
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (postgresql.FindUserByUsernameRow, error) {
	user, err := s.repo.FindUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, derrors.NewErrorf(derrors.ErrorCodeNotFound, "user tidak ditemukan")
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

	pw, _ := hash.HashPassword(data.Password)

	params := postgresql.InsertUserParams{
		ID:        helpers.GenerateID(),
		RoleID:    data.RoleID,
		Username:  data.Username,
		Password:  pw,
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

func (s *Service) UpdatePassword(ctx context.Context, username, oldPassword, newPassword string) error {
	// Get user by username to verify old password
	user, err := s.repo.FindUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return derrors.NewErrorf(derrors.ErrorCodeNotFound, "user tidak ditemukan")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	// Verify old password
	if !hash.CheckPasswordHash(oldPassword, user.Password) {
		return derrors.NewErrorf(derrors.ErrorCodeUnauthorized, "password lama salah")
	}

	// Hash the new password
	hashedPassword, err := hash.HashPassword(newPassword)
	if err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "gagal mengenkripsi password")
	}

	// Update the password
	updateParams := postgresql.UpdatePasswordParams{
		ID:        user.ID,
		Password:  hashedPassword,
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdatePassword(ctx, updateParams)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "gagal memperbarui password")
	}

	return nil
}

func (s *Service) ResetPassword(ctx context.Context, username string) (string, error) {
	// Find user by username first
	user, err := s.repo.FindUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", derrors.NewErrorf(derrors.ErrorCodeNotFound, "user tidak ditemukan")
		}
		return "", derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	// Generate a random password
	newPassword, err := util.GenerateRandomPassword()
	if err != nil {
		return "", derrors.NewErrorf(derrors.ErrorCodeBadRequest, "gagal membuat password acak")
	}

	// Hash the new password
	hashedPassword, err := hash.HashPassword(newPassword)
	if err != nil {
		return "", derrors.NewErrorf(derrors.ErrorCodeBadRequest, "gagal mengenkripsi password")
	}

	// Update the password without checking old password
	updateParams := postgresql.UpdatePasswordParams{
		ID:        user.ID,
		Password:  hashedPassword,
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdatePassword(ctx, updateParams)
	if err != nil {
		return "", derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowsAffected == 0 {
		return "", derrors.NewErrorf(derrors.ErrorCodeNotFound, "gagal memperbarui password")
	}

	// Return the plaintext random password so it can be shared with the user
	return newPassword, nil
}

func (s *Service) DeleteUser(ctx context.Context, username string) error {
	// Find user by username first
	user, err := s.repo.FindUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return derrors.NewErrorf(derrors.ErrorCodeNotFound, "user tidak ditemukan")
		}
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	// Now delete the user using their ID
	params := postgresql.DeleteUserParams{
		ID:        user.ID,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
	rowAffected, err := s.repo.DeleteUser(ctx, params)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	if rowAffected == 0 {
		return derrors.NewErrorf(derrors.ErrorCodeNotFound, "gagal menghapus user")
	}

	return nil
}
