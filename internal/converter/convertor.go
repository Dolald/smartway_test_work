package converter

import (
	"github.com/Dolald/smartway_test_work/internal/domain"
	"github.com/Dolald/smartway_test_work/internal/models"
)

func EmployeeFromModel(emp models.CreateEmployeeRequest) domain.Employee {
	return domain.Employee{
		Name:         emp.Name,
		Surname:      emp.Surname,
		Phone:        emp.Phone,
		Passport:     domain.Passport{Type: emp.Passport.Type, Number: emp.Passport.Number},
		DepartmentId: emp.DepartmentId,
	}
}

func ModelToDomainForUpdate(emp models.UpdateEmployeeRequest) domain.UpdateEmployee {
	var passport domain.UpdatePassport
	if emp.Passport != nil {
		passport = domain.UpdatePassport{
			Type:   emp.Passport.Type,
			Number: emp.Passport.Number,
		}
	}
	return domain.UpdateEmployee{
		Name:         emp.Name,
		Surname:      emp.Surname,
		Phone:        emp.Phone,
		Passport:     &passport,
		DepartmentId: emp.DepartmentId,
	}
}

func DomainToModelEmployee(employees []domain.Employee) []models.EmployeeResponse {
	var employeeRequests []models.EmployeeResponse

	for _, emp := range employees {
		employeeRequest := models.EmployeeResponse{
			ID:           emp.ID,
			DepartmentId: emp.DepartmentId,
			Name:         emp.Name,
			Surname:      emp.Surname,
			Phone:        emp.Phone,
			Passport:     models.Passport(emp.Passport),
		}
		employeeRequests = append(employeeRequests, employeeRequest)
	}

	return employeeRequests
}
