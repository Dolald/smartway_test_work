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
		return 0, fmt.Errorf("failed to create employee: %w", err)
	}

	return input.ID, nil
}

func (r *EmployeeRepository) UpdateEmployee(ctx context.Context, input models.UpdateEmployeeRequest, id int) error {
	var exists bool
	checkQuery := "SELECT EXISTS(SELECT 1 FROM employees WHERE id = $1)"
	err := r.db.QueryRowContext(ctx, checkQuery, id).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if employee exists: %w", err)
	}
	if !exists {
		return fmt.Errorf("employee with id %d does not exist", id)
	}

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
	args = append(args, id)

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *EmployeeRepository) GetEmployeesCompanyDepartment(ctx context.Context, id int) ([]models.EmployeeResponse, error) {
	query := `SELECT * FROM employees WHERE department_id = $1;`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("input failed: %w", err)
	}
	defer rows.Close()

	var list []models.EmployeeResponse

	for rows.Next() {
		var employee models.EmployeeResponse
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Surname, &employee.Phone, &employee.DepartmentId, &employee.Passport.Type, &employee.Passport.Number); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		list = append(list, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	if len(list) == 0 {
		return nil, fmt.Errorf("no employees found for department ID: %d", id)
	}

	return list, nil
}

func (r *EmployeeRepository) GetEmployeesCompany(ctx context.Context, id int) ([]models.EmployeeResponse, error) {

	query := `SELECT e.id, e.name, e.surname, e.phone, e.department_id, e.passport_type, e.passport_number FROM employees e
	          JOIN departments d ON d.id = e.id 
	          WHERE d.company_id = $1`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("input failed: %w", err)
	}
	defer rows.Close()

	var list []models.EmployeeResponse

	for rows.Next() {
		var employee models.EmployeeResponse
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Surname, &employee.Phone, &employee.DepartmentId, &employee.Passport.Type, &employee.Passport.Number); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		list = append(list, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	if len(list) == 0 {
		return nil, fmt.Errorf("no employees found for department ID: %d", id)
	}

	return list, nil
}

func (r *EmployeeRepository) DeleteEmployee(ctx context.Context, id int) error {
	query := `DELETE FROM employees WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("employee delete failed: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no employee found with id: %d", id)
	}

	return nil
}
