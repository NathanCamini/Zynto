package repository

import (
	"Zynto/internal/customers/models"
	"errors"
	"sync"
	"time"
)

type customerRepoMemory struct {
	data map[string]*models.Customer
	mu   sync.RWMutex
}

func NewCustomerRepository() models.CustomerRepository {
	return &customerRepoMemory{
		data: make(map[string]*models.Customer),
	}
}

func (c *customerRepoMemory) CreateCustomer(customer *models.Customer) (*models.Customer, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if customer == nil || customer.ID == "" {
		return nil, errors.New("invalid customer")
	}

	c.data[customer.ID] = customer
	return customer, nil
}

func (c *customerRepoMemory) GetCustomerByID(id string) (*models.Customer, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if customer, ok := c.data[id]; ok {
		return customer, nil
	}

	return nil, errors.New("customer not found")
}

func (c *customerRepoMemory) GetCustomersByCompany(id string) ([]models.Customer, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	customers := make([]models.Customer, 0, len(c.data))
	for _, c := range c.data {
		if c.CompanyId == id {
			customers = append(customers, *c)
		}
	}

	return customers, nil
}

func (c *customerRepoMemory) UpdateCustomer(id string, customer *models.Customer) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[id]; !ok {
		return errors.New("customer not found")
	}

	customer.ID = id
	customer.CreatedAt = c.data[id].CreatedAt
	customer.UpdatedAt = time.Now()

	c.data[id] = customer
	return nil
}

func (c *customerRepoMemory) DeleteCustomer(id string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[id]; !ok {
		return errors.New("customer not found")
	}

	delete(c.data, id)
	return nil
}
