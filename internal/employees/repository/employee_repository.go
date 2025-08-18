package repository

import (
	"Zynto/internal/employees/models"
	"errors"
	"sync"
	"time"
)

type employeeRepoMemory struct {
	data map[string]*models.Employee
	mu   sync.RWMutex
}

func NewEmployeeRepository() models.EmployeeRepository {
	return &employeeRepoMemory{
		data: make(map[string]*models.Employee),
	}
}

func (r *employeeRepoMemory) CreateEmployee(employee *models.Employee) (*models.Employee, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if employee == nil || employee.ID == "" {
		return nil, errors.New("invalid employee")
	}

	r.data[employee.ID] = employee
	return employee, nil
}

func (r *employeeRepoMemory) GetEmployee(id string) (*models.Employee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if employee, ok := r.data[id]; ok {
		return employee, nil
	}

	return nil, errors.New("employee not found")
}

func (r *employeeRepoMemory) GetAllEmployees() ([]models.Employee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	employees := make([]models.Employee, 0, len(r.data))
	for _, e := range r.data {
		employees = append(employees, *e)
	}

	return employees, nil
}

func (r *employeeRepoMemory) UpdateEmployee(id string, employee *models.Employee) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return errors.New("employee not found")
	}

	employee.ID = id
	employee.CreatedAt = r.data[id].CreatedAt
	employee.UpdatedAt = time.Now()

	r.data[id] = employee
	return nil
}

func (r *employeeRepoMemory) DeleteEmployee(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return errors.New("employee not found")
	}

	delete(r.data, id)
	return nil
}
