package models

type EmployeeService interface {
	CreateEmployee(employee *Employee) (*Employee, error)
	UpdateEmployee(id string, employee *Employee) error
	DeleteEmployee(id string) error
	GetEmployeeByID(id string) (*Employee, error)
	GetAllEmployees() ([]Employee, error)
	EmployeeUtils
}

type EmployeeUtils interface {
	EmployeeIsValid(employeeID []string) error
}
