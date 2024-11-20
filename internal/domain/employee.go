package domain

type Passport struct {
	Type   string
	Number string
}

type Employee struct {
	ID           int
	DepartmentId int
	Name         string
	Surname      string
	Phone        string
	Passport     Passport
}

type UpdateEmployee struct {
	DepartmentId int
	Name         *string
	Surname      *string
	Phone        *string
	Passport     *UpdatePassport
}

type UpdatePassport struct {
	Type   *string
	Number *string
}
