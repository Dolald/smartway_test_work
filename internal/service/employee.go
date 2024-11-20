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

func (s *EmployeeService) CreateEmployee(ctx context.Context, employee models.CreateEmployeeRequest) (int, error) {
	employeeId, err := s.repository.CreateEmployee(ctx, converter.ModelToDomainForCteate(employee))
	if err != nil {
		return 0, fmt.Errorf("CreateEmployee failed: %w", err)
	}

	return employeeId, nil
}

func (s *EmployeeService) UpdateEmployee(ctx context.Context, employee models.UpdateEmployeeRequest, id int) error {
	err := s.repository.UpdateEmployee(ctx, converter.ModelToDomainForUpdate(employee), id)
	if err != nil {
		return fmt.Errorf("UpdateEmployee failed: %w", err)
	}

	return nil
}

func (s *EmployeeService) GetEmployeesByDepartmentId(ctx context.Context, id int) ([]models.EmployeeResponse, error) {
	list, err := s.repository.GetEmployeesByDepartmentId(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("GetEmployeesByDepartmentId failed: %w", err)
	}

	responseList := converter.DomainToModelEmployee(list)

	return responseList, nil
}

func (s *EmployeeService) GetEmployeesByCompanyId(ctx context.Context, id int) ([]models.EmployeeResponse, error) {
	list, err := s.repository.GetEmployeesByCompanyId(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("GetEmployeesByCompanyId failed: %w", err)
	}

	responseList := converter.DomainToModelEmployee(list)

	return responseList, nil
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id int) error {
	err := s.repository.DeleteEmployee(ctx, id)
	if err != nil {
		return fmt.Errorf("DeleteEmployee failed: %w", err)
	}

	return nil
}
