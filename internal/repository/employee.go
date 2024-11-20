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

func (r *EmployeeRepository) CreateEmployee(ctx context.Context, input domain.Employee) (int, error) {
	query := `INSERT INTO employees (name, surname, phone, department_id, passport_type, passport_number) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

	var employeeId int

	err := r.db.QueryRowContext(ctx, query, input.Name, input.Surname, input.Phone, input.DepartmentId, input.Passport.Type, input.Passport.Number).Scan(&employeeId)
	if err != nil {
		return 0, fmt.Errorf("failed to scan input.ID: %w", err)
	}

	return employeeId, nil
}

func (r *EmployeeRepository) UpdateEmployee(ctx context.Context, input domain.UpdateEmployee, id int) error {
	values := make([]string, 0)
	args := make([]any, 0)
	argsId := 1

	if input.Name != nil {
		values = append(values, fmt.Sprintf("name=$%d", argsId))
		args = append(args, input.Name)
		argsId++
	}

	if input.Surname != nil {
		values = append(values, fmt.Sprintf("surname=$%d", argsId))
		args = append(args, input.Surname)
		argsId++
	}

	if input.Phone != nil {
		values = append(values, fmt.Sprintf("phone=$%d", argsId))
		args = append(args, input.Phone)
		argsId++
	}

	if input.DepartmentId != 0 {
		values = append(values, fmt.Sprintf("department_id=$%d", argsId))
		args = append(args, input.DepartmentId)
		argsId++
	}

	if input.Passport.Type != nil {
		values = append(values, fmt.Sprintf("passport_type=$%d", argsId))
		args = append(args, input.Passport.Type)
		argsId++
	}

	if input.Passport.Number != nil {
		values = append(values, fmt.Sprintf("passport_number=$%d", argsId))
		args = append(args, input.Passport.Number)
		argsId++
	}

	setQuery := strings.Join(values, ", ")

	query := fmt.Sprintf("UPDATE employees SET %s WHERE id = $%d", setQuery, argsId)
	args = append(args, id)

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("ExecContext failed: %w", err)
	}

	return nil
}

func (r *EmployeeRepository) GetEmployeesByDepartmentId(ctx context.Context, id int) ([]domain.Employee, error) {
	query := `SELECT * FROM employees WHERE department_id = $1;`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("input failed: %w", err)
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

func (r *EmployeeRepository) GetEmployeesCompany(ctx context.Context, id int) ([]domain.Employee, error) {

	query := `SELECT e.id, e.name, e.surname, e.phone, e.department_id, e.passport_type, e.passport_number FROM employees e
	          JOIN departments d ON d.id = e.id 
	          WHERE d.company_id = $1`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("input failed: %w", err)
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
