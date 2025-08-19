package main

import (
	cus_cont "Zynto/internal/customers/controllers"
	cus_repo "Zynto/internal/customers/repository"
	cus_serv "Zynto/internal/customers/service"
	emp_cont "Zynto/internal/employees/controllers"
	emp_repo "Zynto/internal/employees/repository"
	emp_serv "Zynto/internal/employees/service"
	ser_cont "Zynto/internal/services/controllers"
	ser_repo "Zynto/internal/services/repository"
	ser_serv "Zynto/internal/services/service"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}
	slog.Info("all systems offline")
}

func run() error {
	handler := chi.NewMux()
	handler.Use(middleware.Recoverer)
	handler.Use(middleware.RequestID)
	handler.Use(middleware.Logger)

	employeeRepo := emp_repo.NewEmployeeRepository()
	employeeService := emp_serv.NewEmployeeService(employeeRepo)
	employeesController := emp_cont.NewEmployeesController(employeeService)

	serviceRepo := ser_repo.NewServiceRepository()
	serviceService := ser_serv.NewServiceService(serviceRepo, employeeService)
	servicesController := ser_cont.NewServicesController(serviceService)

	constumerRepo := cus_repo.NewCustomerRepository()
	customerService := cus_serv.NewCustomerService(constumerRepo)
	customersController := cus_cont.NewCustomersController(customerService)

	employeesController.RegisterRoutes(handler)
	servicesController.RegisterRoutes(handler)
	customersController.RegisterRoutes(handler)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	fmt.Println("✅ Servidor rodando em http://localhost:8080 🚀")

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
