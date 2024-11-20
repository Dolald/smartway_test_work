package repository

import (
	"context"

	"github.com/Dolald/smartway_test_work/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Employee interface {
	CreateEmployee(ctx context.Context, input domain.Employee) (int, error)
	UpdateEmployee(ctx context.Context, input domain.UpdateEmployee, id int) error
	DeleteEmployee(ctx context.Context, id int) error
	GetEmployeesByDepartmentId(ctx context.Context, id int) ([]domain.Employee, error)
	GetEmployeesCompany(ctx context.Context, id int) ([]domain.Employee, error)
}

type Repository struct {
	Employee
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Employee: NewEmployeeRepository(db),
	}
}
