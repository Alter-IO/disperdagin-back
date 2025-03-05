package service

import (
	"alter-io-go/repositories/postgresql"
)

const postgreErrMsg = "PostgreSQL Error"

type Service struct {
	repo *postgresql.Queries
}

func NewService(repo *postgresql.Queries) *Service {
	return &Service{repo}
}
