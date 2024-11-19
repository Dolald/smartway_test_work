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
	CompanyID  int        `json:"company_id"`
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Phone      string     `json:"phone"`
	Passport   Passport   `json:"passport"`
	Department Department `json:"department"`
}

type EmployeeResponse struct {
	ID           int
	DepartmentId int
	Name         string
	Surname      string
	Phone        string
	Passport     Passport
}

type EmployeeRequest struct {
	ID           int      `json:"-"`
	DepartmentId int      `json:"department_id" binding:"required"`
	Name         string   `json:"name" binding:"required"`
	Surname      string   `json:"surname" binding:"required"`
	Phone        string   `json:"phone" binding:"required"`
	Passport     Passport `json:"passport" binding:"required"`
}

type UpdateEmployeeRequest struct {
	DepartmentId int            `json:"department_id"`
	Name         *string        `json:"name"`
	Surname      *string        `json:"surname"`
	Phone        *string        `json:"phone"`
	Passport     UpdatePassport `json:"passport"`
}

type UpdatePassport struct {
	Type   *string `json:"type"`
	Number *string `json:"number"`
}
