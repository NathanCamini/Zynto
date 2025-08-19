package service

import (
	"Zynto/internal/customers/models"
	"errors"
	"time"

	"github.com/google/uuid"
)

type CustomerService struct {
	customerRepository models.CustomerRepository
}

func NewCustomerService(repo models.CustomerRepository) *CustomerService {
	return &CustomerService{customerRepository: repo}
}

func (c *CustomerService) CreateCustomer(customer *models.Customer) (*models.Customer, error) {
	if customer == nil {
		return nil, errors.New("customer is nil")
	}

	genderValid := customer.Gender.IsValid()
	if !genderValid {
		return nil, errors.New("gender is not valid")
	}

	customer.ID = uuid.New().String()
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	return c.customerRepository.CreateCustomer(customer)
}

func (c *CustomerService) GetCustomerByID(id string) (*models.Customer, error) {
	if id == "" {
		return nil, errors.New("customer ID is empty")
	}

	customer, err := c.customerRepository.GetCustomerByID(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *CustomerService) GetCustomersByCompany(id string) ([]models.Customer, error) {
	if id == "" {
		return nil, errors.New("customer ID is empty")
	}

	customers, err := c.customerRepository.GetCustomersByCompany(id)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (c *CustomerService) UpdateCustomer(id string, customer *models.Customer) error {
	if id == "" {
		return errors.New("customer ID is empty")
	}

	err := c.customerRepository.UpdateCustomer(id, customer)
	if err != nil {
		return err
	}

	return nil
}

func (c *CustomerService) DeleteCustomer(id string) error {
	if id == "" {
		return errors.New("customer ID is empty")
	}

	return c.customerRepository.DeleteCustomer(id)
}
