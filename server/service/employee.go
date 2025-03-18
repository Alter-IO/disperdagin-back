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

// GetAllEmployees retrieves all employees
func (s *Service) GetAllEmployees(ctx context.Context) ([]postgresql.FindAllEmployeesRow, error) {
	employees, err := s.repo.FindAllEmployees(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return employees, nil
}

// GetActiveEmployees retrieves only active employees
func (s *Service) GetActiveEmployees(ctx context.Context) ([]postgresql.FindActiveEmployeesRow, error) {
	employees, err := s.repo.FindActiveEmployees(ctx)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return employees, nil
}

// GetEmployeesByPosition retrieves employees filtered by position
func (s *Service) GetEmployeesByPosition(ctx context.Context, position string) ([]postgresql.FindEmployeesByPositionRow, error) {
	employees, err := s.repo.FindEmployeesByPosition(ctx, position)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return employees, nil
}

// GetEmployeeByID retrieves a single employee by ID
func (s *Service) GetEmployeeByID(ctx context.Context, id string) (postgresql.FindEmployeeByIDRow, error) {
	employee, err := s.repo.FindEmployeeByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return employee, derrors.NewErrorf(derrors.ErrorCodeNotFound, "karyawan tidak ditemukan")
		}
		return employee, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return employee, nil
}

func validateEmployee(data postgresql.InsertEmployeeParams) error {
	if data.Name == "" {
		return errors.New("nama karyawan wajib diisi")
	}

	if data.Position == "" {
		return errors.New("jabatan karyawan wajib diisi")
	}

	return nil
}

// CreateEmployee creates a new employee
func (s *Service) CreateEmployee(ctx context.Context, data postgresql.InsertEmployeeParams) error {
	if err := validateEmployee(data); err != nil {
		return derrors.NewErrorf(derrors.ErrorCodeBadRequest, "%s", err.Error())
	}

	// Generate a new ID
	params := postgresql.InsertEmployeeParams{
		ID:         helpers.GenerateID(),
		Name:       data.Name,
		Position:   data.Position,
		Address:    data.Address,
		EmployeeID: data.EmployeeID,
		Birthplace: data.Birthplace,
		Birthdate:  data.Birthdate,
		Photo:      data.Photo,
		Status:     data.Status,
		Author:     data.Author,
		CreatedAt:  pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	if err := s.repo.InsertEmployee(ctx, params); err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return nil
}

// UpdateEmployee updates an existing employee
func (s *Service) UpdateEmployee(ctx context.Context, data postgresql.UpdateEmployeeParams) (int64, error) {
	// Validate data
	if data.Name == "" {
		return 0, derrors.NewErrorf(derrors.ErrorCodeBadRequest, "nama karyawan wajib diisi")
	}

	if data.Position == "" {
		return 0, derrors.NewErrorf(derrors.ErrorCodeBadRequest, "jabatan karyawan wajib diisi")
	}

	// Add current timestamp
	params := postgresql.UpdateEmployeeParams{
		ID:         data.ID,
		Name:       data.Name,
		Position:   data.Position,
		Address:    data.Address,
		EmployeeID: data.EmployeeID,
		Birthplace: data.Birthplace,
		Birthdate:  data.Birthdate,
		Photo:      data.Photo,
		Status:     data.Status,
		UpdatedAt:  pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.UpdateEmployee(ctx, params)
	if err != nil {
		return 0, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return rowsAffected, nil
}

// DeleteEmployee soft deletes an employee
func (s *Service) DeleteEmployee(ctx context.Context, id string) (int64, error) {
	params := postgresql.DeleteEmployeeParams{
		ID:        id,
		DeletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rowsAffected, err := s.repo.DeleteEmployee(ctx, params)
	if err != nil {
		return 0, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, postgreErrMsg)
	}

	return rowsAffected, nil
}
