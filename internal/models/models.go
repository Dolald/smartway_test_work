package models

type Passport struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Department struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Employee struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Phone      string     `json:"phone"`
	CompanyID  int        `json:"company_id"`
	Passport   Passport   `json:"passport"`
	Department Department `json:"department"`
}

type EmployeesListResponse struct {
	Employees []Employee
}

type EmployeeRequest struct {
	ID           int      `json:"-"`
	Name         string   `json:"name" binding:"required"`
	Surname      string   `json:"surname" binding:"required"`
	Phone        string   `json:"phone" binding:"required"`
	Passport     Passport `json:"passport" binding:"required"`
	DepartmentId int      `json:"department_id" binding:"required"`
}

type UpdatedEmployeeRequest struct {
	ID           int             `json:"id"`
	Name         *string         `json:"name, omitempty"`
	Surname      *string         `json:"surname, omitempty"`
	Phone        *string         `json:"phone, omitempty"`
	Passport     UpdatedPassport `json:"passport, omitempty"`
	DepartmentId int             `json:"department_id"`
}

type UpdatedPassport struct {
	Type   *string `json:"type, omitempty"`
	Number *string `json:"number, omitempty"`
}
