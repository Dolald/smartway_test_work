package service

import (
	"context"

	"github.com/Dolald/smartway_test_work/internal/models"
	"github.com/Dolald/smartway_test_work/internal/repository"
)

type Employee interface {
	CreateEmployee(ctx context.Context, employee models.CreateEmployeeRequest) (int, error)
	UpdateEmployee(ctx context.Context, employee models.UpdateEmployeeRequest, id int) error
	GetEmployeesByDepartmentId(ctx context.Context, id int) ([]models.EmployeeResponse, error)
	GetEmployeesByCompanyId(ctx context.Context, id int) ([]models.EmployeeResponse, error)
	DeleteEmployee(ctx context.Context, id int) error
}

type Service struct {
	Employee
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Employee: NewEmployeeService(repository.Employee),
	}
}
