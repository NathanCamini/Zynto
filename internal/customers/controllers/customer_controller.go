package controllers

import (
	"Zynto/internal/customers/models"
	"Zynto/internal/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type CustomersController struct {
	customerService models.CustomerService
}

func NewCustomersController(customerService models.CustomerService) *CustomersController {
	return &CustomersController{customerService: customerService}
}

func (c *CustomersController) RegisterRoutes(r chi.Router) {
	r.Route("/customers", func(r chi.Router) {
		r.Post("/", c.createCustomer())
		r.Get("/{customerID}", c.getCustomerByID())
		r.Get("/allCustomers/{companyID}", c.getCustomersByCompany())
		r.Put("/{customerID}", c.updateCustomer())
		r.Delete("/{customerID}", c.deleteCustomer())
	})
}

func (c *CustomersController) createCustomer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer models.Customer

		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			utils.SendJson(w, utils.Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		createdCustomer, err := c.customerService.CreateCustomer(&customer)
		if err != nil {
			utils.SendJson(w, utils.Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		utils.SendJson(w, utils.Response{Data: createdCustomer}, http.StatusCreated)
	}
}

func (c *CustomersController) getCustomerByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerID := chi.URLParam(r, "customerID")

		customer, err := c.customerService.GetCustomerByID(customerID)
		if err != nil {
			utils.SendJson(w, utils.Response{Error: err.Error()}, http.StatusNotFound)
			return
		}
		utils.SendJson(w, utils.Response{Data: customer}, http.StatusOK)
	}
}

func (c *CustomersController) getCustomersByCompany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		companyID := chi.URLParam(r, "companyID")

		customers, err := c.customerService.GetCustomersByCompany(companyID)
		if err != nil {
			utils.SendJson(w, utils.Response{Error: err.Error()}, http.StatusInternalServerError)
			return
		}

		utils.SendJson(w, utils.Response{Data: customers}, http.StatusOK)
	}
}

func (c *CustomersController) updateCustomer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerID := chi.URLParam(r, "customerID")
		var customer models.Customer

		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			utils.SendJson(w, utils.Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if err := c.customerService.UpdateCustomer(customerID, &customer); err != nil {
			utils.SendJson(w, utils.Response{Error: err.Error()}, http.StatusNotFound)
			return
		}

		utils.SendJson(w, utils.Response{Data: customer}, http.StatusOK)
	}
}

func (c *CustomersController) deleteCustomer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerID := chi.URLParam(r, "customerID")

		if err := c.customerService.DeleteCustomer(customerID); err != nil {
			utils.SendJson(w, utils.Response{Error: err.Error()}, http.StatusNotFound)
			return
		}

		utils.SendJson(w, utils.Response{Data: "customer deleted successfully"}, http.StatusOK)
	}
}
