package repository

import (
	"Zynto/internal/services/models"
	"errors"
	"sync"
	"time"
)

type serviceRepoMemory struct {
	data map[string]*models.Service
	mu   sync.RWMutex
}

func NewServiceRepository() models.ServiceRepository {
	return &serviceRepoMemory{
		data: make(map[string]*models.Service),
	}
}

func (r *serviceRepoMemory) CreateService(service *models.Service) (*models.Service, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if service == nil || service.ID == "" {
		return nil, errors.New("invalid service")
	}

	r.data[service.ID] = service
	return service, nil
}

func (r *serviceRepoMemory) GetService(id string) (*models.Service, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if service, ok := r.data[id]; ok {
		return service, nil
	}

	return nil, errors.New("service not found")
}

func (r *serviceRepoMemory) GetAllServices() ([]models.Service, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	services := make([]models.Service, 0, len(r.data))
	for _, s := range r.data {
		services = append(services, *s)
	}

	return services, nil
}

func (r *serviceRepoMemory) UpdateService(id string, service *models.Service) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return errors.New("service not found")
	}

	service.ID = id
	service.CreatedAt = r.data[id].CreatedAt
	service.UpdatedAt = time.Now()

	r.data[id] = service
	return nil
}

func (r *serviceRepoMemory) DeleteService(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return errors.New("service not found")
	}
	
	delete(r.data, id)
	return nil
}
