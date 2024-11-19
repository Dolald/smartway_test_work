package service

import (
	"context"
	"fmt"

	"github.com/Dolald/smartway_test_work/internal/models"
	"github.com/Dolald/smartway_test_work/internal/repository"
)

type EmployeeService struct {
	repository repository.Employee
}

func NewEmployeeService(repository repository.Employee) *EmployeeService {
	return &EmployeeService{repository: repository}
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, input models.EmployeeRequest) (int, error) {
	workerId, err := s.repository.CreateEmployee(ctx, input)
	if err != nil {
		return 0, fmt.Errorf("CreateEmployee failed: %w", err)
	}

	return workerId, nil
}

func (s *EmployeeService) UpdateEmployee(ctx context.Context, input models.UpdateEmployeeRequest) error {
	err := s.repository.UpdateEmployee(ctx, input)
	if err != nil {
		return fmt.Errorf("UpdateEmployee failed: %w", err)
	}

	return nil
}

func (s *EmployeeService) GetEmployeesCompanyDepartment(ctx context.Context, id int) ([]models.EmployeeResponse, error) {
	list, err := s.repository.GetEmployeesCompanyDepartment(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("GetEmployeesCompanyDepartment failed: %w", err)
	}

	return list, nil
}

func (s *EmployeeService) GetEmployeesCompany(ctx context.Context, id int) ([]models.EmployeeResponse, error) {
	list, err := s.repository.GetEmployeesCompany(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("GetEmployeesCompany failed: %w", err)
	}

	return list, nil
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id int) error {
	err := s.repository.DeleteEmployee(ctx, id)
	if err != nil {
		return fmt.Errorf("DeleteEmployee failed: %w", err)
	}

	return nil
}
