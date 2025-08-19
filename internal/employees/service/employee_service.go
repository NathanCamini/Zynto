package service

import (
	"Zynto/internal/employees/models"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

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

	genderValid := employee.Gender.IsValid()
	if !genderValid {
		return nil, errors.New("gender is not valid")
	}

	employee.ID = uuid.New().String()
	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()

	return e.employeeRepository.CreateEmployee(employee)
}

func (e *EmployeeService) GetEmployeeByID(id string) (*models.Employee, error) {
	if id == "" {
		return nil, errors.New("employee ID is empty")
	}

	employee, err := e.employeeRepository.GetEmployeeByID(id)
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

func (e *EmployeeService) EmployeeIsValid(employeeID []string) error {
	for _, employeeID := range employeeID {
		_, err := e.GetEmployeeByID(employeeID)
		if err != nil {
			return fmt.Errorf("failed to get employee %s: %w", employeeID, err)
		}
	}

	return nil
}
