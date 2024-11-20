package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Dolald/smartway_test_work/internal/domain"
	"github.com/jmoiron/sqlx"
)

type EmployeeRepository struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) CreateEmployee(ctx context.Context, employee domain.Employee) (int, error) {
	query := `INSERT INTO employees (name, surname, phone, department_id, passport_type, passport_number) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

	var employeeId int

	err := r.db.QueryRowContext(ctx, query, employee.Name, employee.Surname, employee.Phone, employee.DepartmentId, employee.Passport.Type, employee.Passport.Number).Scan(&employeeId)
	if err != nil {
		return 0, fmt.Errorf("failed to scan employee.ID: %w", err)
	}

	return employeeId, nil
}

func (r *EmployeeRepository) UpdateEmployee(ctx context.Context, employee domain.UpdateEmployee, id int) error {
	values := make([]string, 0)
	args := make([]any, 0)
	argsId := 1

	if employee.Name != nil {
		values = append(values, fmt.Sprintf("name=$%d", argsId))
		args = append(args, employee.Name)
		argsId++
	}

	if employee.Surname != nil {
		values = append(values, fmt.Sprintf("surname=$%d", argsId))
		args = append(args, employee.Surname)
		argsId++
	}

	if employee.Phone != nil {
		values = append(values, fmt.Sprintf("phone=$%d", argsId))
		args = append(args, employee.Phone)
		argsId++
	}

	if employee.DepartmentId != 0 {
		values = append(values, fmt.Sprintf("department_id=$%d", argsId))
		args = append(args, employee.DepartmentId)
		argsId++
	}

	if employee.Passport.Type != nil {
		values = append(values, fmt.Sprintf("passport_type=$%d", argsId))
		args = append(args, employee.Passport.Type)
		argsId++
	}

	if employee.Passport.Number != nil {
		values = append(values, fmt.Sprintf("passport_number=$%d", argsId))
		args = append(args, employee.Passport.Number)
		argsId++
	}

	setQuery := strings.Join(values, ", ")

	query := fmt.Sprintf("UPDATE employees SET %s WHERE id = $%d", setQuery, argsId)
	args = append(args, id)

	rows, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("ExecContext failed: %w", err)
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("employee with id %d not found", id)
	}

	return nil
}

func (r *EmployeeRepository) GetEmployeesByDepartmentId(ctx context.Context, id int) ([]domain.Employee, error) {
	query := `SELECT * FROM employees WHERE department_id = $1;`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("employee failed: %w", err)
	}
	defer rows.Close()

	var list []domain.Employee

	for rows.Next() {
		var employee domain.Employee
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Surname, &employee.Phone, &employee.DepartmentId, &employee.Passport.Type, &employee.Passport.Number); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		list = append(list, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return list, nil
}

func (r *EmployeeRepository) GetEmployeesByCompanyId(ctx context.Context, id int) ([]domain.Employee, error) {

	query := `SELECT e.id, e.name, e.surname, e.phone, e.department_id, e.passport_type, e.passport_number FROM employees e
	          JOIN departments d ON d.id = e.id 
	          WHERE d.company_id = $1`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("employee failed: %w", err)
	}
	defer rows.Close()

	var list []domain.Employee

	for rows.Next() {
		var employee domain.Employee
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Surname, &employee.Phone, &employee.DepartmentId, &employee.Passport.Type, &employee.Passport.Number); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		list = append(list, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
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
