package service

import (
	"alter-io-go/domain"
	"alter-io-go/helpers/derrors"
	"alter-io-go/helpers/hash"
	"alter-io-go/helpers/jwt"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (s *Service) Login(ctx context.Context, data domain.LoginReq) (domain.LoginResp, error) {
	if err := data.Validate(); err != nil {
		return domain.LoginResp{}, derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	user, err := s.repo.FindUserByUsername(ctx, data.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.LoginResp{}, derrors.NewErrorf(derrors.ErrorCodeNotFound, "username atau password salah")
		}

		return domain.LoginResp{}, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, "%s", err.Error())
	}

	if !hash.CheckPasswordHash(data.Password, user.Password) {
		return domain.LoginResp{}, derrors.NewErrorf(derrors.ErrorCodeNotFound, "username atau password salah")
	}

	tokenStr, err := jwt.GenerateJWT(user.ID, user.RoleID)
	if err != nil {
		return domain.LoginResp{}, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, "failed to generate token for user id: %s and error: %s", user.ID, err.Error())
	}

	authData := domain.LoginResp{
		ID:          user.ID,
		RoleID:      user.RoleID,
		Username:    data.Username,
		AccessToken: tokenStr,
	}
	return authData, nil
}
