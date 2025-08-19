package models

type ServiceService interface {
	CreateService(service *Service) (*Service, error)
	UpdateService(id string, service *Service) error
	DeleteService(id string) error
	GetServiceByID(id string) (*Service, error)
	GetAllServices() ([]Service, error)
}
