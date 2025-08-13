package service

import (
	"Zynto/internal/employees/models"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Aqui devo criar o UUID e depois chamar a função do repository para inserir no BD
// Deve ter as regras de negocios apenas aqui
type EmployeeService struct {
	employeeRepository models.EmployeeRepository
}

func NewEmployeeService(repo models.EmployeeRepository) *EmployeeService {
	return &EmployeeService{employeeRepository: repo}
}

func (e *EmployeeService) CreateEmployee(employee *models.Employee) (*models.Employee, error) {
	if employee == nil {
		return nil, errors.New("employee is nil")
	}
	employee.ID = uuid.New().String()
	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()

	return e.employeeRepository.CreateEmployee(employee)
}

func (e *EmployeeService) GetEmployee(id string) (*models.Employee, error) {
	if id == "" {
		return nil, errors.New("employee ID is empty")
	}
	employee, err := e.employeeRepository.GetEmployee(id)
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (e *EmployeeService) GetAllEmployees() ([]models.Employee, error) {
	employees, err := e.employeeRepository.GetAllEmployees()
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (e *EmployeeService) UpdateEmployee(id string, employee *models.Employee) error {
	if id == "" {
		return errors.New("employee ID is empty")
	}

	err := e.employeeRepository.UpdateEmployee(id, employee)
	if err != nil {
		return err
	}

	return nil
}

func (e *EmployeeService) DeleteEmployee(id string) error {
	if id == "" {
		return errors.New("employee ID is empty")
	}

	return e.employeeRepository.DeleteEmployee(id)
}
