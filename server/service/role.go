package service

import (
	"alter-io-go/helpers/derrors"
	"alter-io-go/repositories/postgresql"
	"context"
)

func (s *Service) GetRoles(ctx context.Context) ([]postgresql.FindRolesRow, error) {
	roles, err := s.repo.FindRoles(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return roles, nil
}
