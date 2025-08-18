package service

import (
	employees "Zynto/internal/employees/service"
	"Zynto/internal/services/models"
	"errors"
	"time"

	"github.com/google/uuid"
)

type ServiceService struct {
	serviceRepository models.ServiceRepository
	employeeService   *employees.EmployeeService
}

func NewServiceService(repo models.ServiceRepository, employeeService *employees.EmployeeService) *ServiceService {
	return &ServiceService{
		serviceRepository: repo,
		employeeService:   employeeService,
	}
}

func (s *ServiceService) CreateService(service *models.Service) (*models.Service, error) {
	if service == nil {
		return nil, errors.New("service is nil")
	}

	genderServiceValid := service.GenderService.IsValid()
	if !genderServiceValid {
		return nil, errors.New("gender is not valid")
	}

	err := s.employeeService.EmployeeIsValid(service.EmployeeId)
	if err != nil {
		return nil, err
	}

	service.ID = uuid.New().String()
	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()

	return s.serviceRepository.CreateService(service)
}

func (s *ServiceService) GetService(id string) (*models.Service, error) {
	if id == "" {
		return nil, errors.New("service ID is empty")
	}

	service, err := s.serviceRepository.GetService(id)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (s *ServiceService) GetAllServices() ([]models.Service, error) {
	services, err := s.serviceRepository.GetAllServices()
	if err != nil {
		return nil, err
	}

	return services, nil
}

func (s *ServiceService) UpdateService(id string, service *models.Service) error {
	if id == "" {
		return errors.New("service ID is empty")
	}

	return s.serviceRepository.UpdateService(id, service)
}

func (s *ServiceService) DeleteService(id string) error {
	if id == "" {
		return errors.New("service ID is empty")
	}
	return s.serviceRepository.DeleteService(id)
}
