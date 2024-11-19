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

type EmployeeResponse struct {
	ID           int
	Name         string
	Surname      string
	Phone        string
	DepartmentId int
	Passport     Passport
}

type EmployeeRequest struct {
	ID           int      `json:"-"`
	Name         string   `json:"name" binding:"required"`
	Surname      string   `json:"surname" binding:"required"`
	Phone        string   `json:"phone" binding:"required"`
	Passport     Passport `json:"passport" binding:"required"`
	DepartmentId int      `json:"department_id" binding:"required"`
}

type UpdateEmployeeRequest struct {
	ID           int            `json:"id"`
	Name         *string        `json:"name"`
	Surname      *string        `json:"surname"`
	Phone        *string        `json:"phone"`
	Passport     UpdatePassport `json:"passport"`
	DepartmentId int            `json:"department_id"`
}

type UpdatePassport struct {
	Type   *string `json:"type"`
	Number *string `json:"number"`
}
