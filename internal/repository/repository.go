package repository

import (
	"context"

	"github.com/Dolald/smartway_test_work/internal/models"
	"github.com/jmoiron/sqlx"
)

type Employee interface {
	CreateEmployee(ctx context.Context, input models.EmployeeRequest) (int, error)
	UpdateEmployee(ctx context.Context, input models.UpdateEmployeeRequest, id int) error
	DeleteEmployee(ctx context.Context, id int) error
	GetEmployeesCompanyDepartment(ctx context.Context, id int) ([]models.EmployeeResponse, error)
	GetEmployeesCompany(ctx context.Context, id int) ([]models.EmployeeResponse, error)
}

type Repository struct {
	Employee
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Employee: NewEmployeeRepository(db),
	}
}
