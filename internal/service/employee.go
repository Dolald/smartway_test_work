package service

import (
	"context"
	"fmt"

	"github.com/Dolald/smartway_test_work/internal/converter"
	"github.com/Dolald/smartway_test_work/internal/models"
	"github.com/Dolald/smartway_test_work/internal/repository"
)

type EmployeeService struct {
	repository repository.Employee
}

func NewEmployeeService(repository repository.Employee) *EmployeeService {
	return &EmployeeService{repository: repository}
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, input models.CreateEmployeeRequest) (int, error) {
	workerId, err := s.repository.CreateEmployee(ctx, converter.ModelToDomainForCteate(input))
	if err != nil {
		return 0, fmt.Errorf("CreateEmployee failed: %w", err)
	}

	return workerId, nil
}

func (s *EmployeeService) UpdateEmployee(ctx context.Context, input models.UpdateEmployeeRequest, id int) error {
	err := s.repository.UpdateEmployee(ctx, converter.ModelToDomainForUpdate(input), id)
	if err != nil {
		return fmt.Errorf("UpdateEmployee failed: %w", err)
	}

	return nil
}

func (s *EmployeeService) GetEmployeesByDepartmentId(ctx context.Context, id int) ([]models.EmployeeResponse, error) {
	// var exists bool
	// checkQuery := "SELECT COUNT(*) FROM employees WHERE id = $1)"
	// err := r.db.QueryRowContext(ctx, checkQuery, id).Scan(&exists)
	// if err != nil {
	// 	return fmt.Errorf("failed to check if employee exists: %w", err)
	// }
	// if !exists {
	// 	return fmt.Errorf("employee with id %d does not exist", id)
	// }

	list, err := s.repository.GetEmployeesByDepartmentId(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("GetEmployeesByDepartmentId failed: %w", err)
	}

	responseList := converter.DomainToModel(list)

	return responseList, nil
}

func (s *EmployeeService) GetEmployeesCompany(ctx context.Context, id int) ([]models.EmployeeResponse, error) {
	list, err := s.repository.GetEmployeesCompany(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("GetEmployeesCompany failed: %w", err)
	}

	responseList := converter.DomainToModel(list)

	return responseList, nil
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id int) error {
	err := s.repository.DeleteEmployee(ctx, id)
	if err != nil {
		return fmt.Errorf("DeleteEmployee failed: %w", err)
	}

	return nil
}
