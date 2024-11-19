package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Dolald/smartway_test_work/internal/models"
	"github.com/jmoiron/sqlx"
)

type EmployeeRepository struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) CreateEmployee(ctx context.Context, input models.EmployeeRequest) (int, error) {
	query := `INSERT INTO employees (name, surname, phone, department_id, passport_type, passport_number) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

	err := r.db.QueryRowContext(ctx, query, input.Name, input.Surname, input.Phone, input.DepartmentId, input.Passport.Type, input.Passport.Number).Scan(&input.ID)
	if err != nil {
		return 0, fmt.Errorf("input failed: %w", err)
	}

	return input.ID, nil
}

func (r *EmployeeRepository) UpdateEmployee(ctx context.Context, input models.UpdatedEmployeeRequest) error {
	values := make([]string, 0)
	args := make([]any, 0)
	argsId := 1

	if input.Name != nil {
		values = append(values, fmt.Sprintf("name=$%d", argsId))
		args = append(args, *input.Name)
		argsId++
	}

	if input.Surname != nil {
		values = append(values, fmt.Sprintf("surname=$%d", argsId))
		args = append(args, *input.Surname)
		argsId++
	}

	if input.Phone != nil {
		values = append(values, fmt.Sprintf("phone=$%d", argsId))
		args = append(args, *input.Phone)
		argsId++
	}

	if input.DepartmentId != 0 {
		values = append(values, fmt.Sprintf("department_id=$%d", argsId))
		args = append(args, input.DepartmentId)
		argsId++
	}

	if input.Passport.Type != nil {
		values = append(values, fmt.Sprintf("passport_type=$%d", argsId))
		args = append(args, *input.Passport.Type)
		argsId++
	}

	if input.Passport.Number != nil {
		values = append(values, fmt.Sprintf("passport_number=$%d", argsId))
		args = append(args, *input.Passport.Number)
		argsId++
	}

	setQuery := strings.Join(values, ", ")

	query := fmt.Sprintf("UPDATE employees SET %s WHERE id = $%d", setQuery, argsId)
	args = append(args, input.ID)

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *EmployeeRepository) getWorkersCompanyDepartment(ctx context.Context, id int) (models.EmployeesListResponse, error) {

	return models.EmployeesListResponse{}, nil
}
