package models

type CustomerRepository interface {
	CreateCustomer(customer *Customer) (*Customer, error)
	UpdateCustomer(id string, customer *Customer) error
	DeleteCustomer(id string) error
	GetCustomerByID(id string) (*Customer, error)
	GetCustomersByCompany(companyId string) ([]Customer, error)
}
