package service

import (
	"context"
	"fmt"

	"github.com/Dolald/smartway_test_work/internal/models"
	"github.com/Dolald/smartway_test_work/internal/repository"
)

type WorkerService struct {
	repository repository.Employee
}

func NewEmployeeService(repository repository.Employee) *WorkerService {
	return &WorkerService{repository: repository}
}

func (s *WorkerService) CreateEmployee(ctx context.Context, input models.EmployeeRequest) (int, error) {
	workerId, err := s.repository.CreateEmployee(ctx, input)
	if err != nil {
		return 0, fmt.Errorf("CreateWorker failed: %w", err)
	}

	return workerId, nil
}

func (s *WorkerService) UpdateEmployee(ctx context.Context, input models.UpdatedEmployeeRequest) error {
	err := s.repository.UpdateEmployee(ctx, input)
	if err != nil {
		return fmt.Errorf("UpdateEmployee failed: %w", err)
	}

	return nil
}

func (r *WorkerService) GetEmployeesCompanyDepartment(ctx context.Context, id int) (models.EmployeesListResponse, error) {

	return models.EmployeesListResponse{}, nil
}
