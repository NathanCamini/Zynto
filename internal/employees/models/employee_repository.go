package models

type EmployeeRepository interface {
	CreateEmployee(employee *Employee) (*Employee, error)
	GetEmployee(id string) (*Employee, error)
	GetAllEmployees() ([]Employee, error)
	UpdateEmployee(id string, employee *Employee) error
	DeleteEmployee(id string) error
}
