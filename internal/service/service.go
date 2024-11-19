package service

import (
	"context"

	"github.com/Dolald/smartway_test_work/internal/models"
	"github.com/Dolald/smartway_test_work/internal/repository"
)

type Employee interface {
	CreateEmployee(ctx context.Context, input models.EmployeeRequest) (int, error)
	UpdateEmployee(ctx context.Context, input models.UpdatedEmployeeRequest) error
	GetEmployeesCompanyDepartment(ctx context.Context, id int) (models.EmployeesListResponse, error)
}

type Service struct {
	Employee
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Employee: NewEmployeeService(repository.Employee),
	}
}