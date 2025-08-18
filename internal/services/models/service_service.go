package models

type ServiceService interface {
	CreateService(service *Service) (*Service, error)
	GetService(id string) (*Service, error)
	GetAllServices() ([]Service, error)
	UpdateService(id string, service *Service) error
	DeleteService(id string) error
}
