package dependency_injection

import (
	controller "Zynto/internal/employees/controllers"
	repository "Zynto/internal/employees/repository"
	service "Zynto/internal/employees/service"
)

func NewAppEmployeesController() *controller.EmployeesController {
	return controller.NewEmployeesController(
		service.NewEmployeeService(
			repository.NewEmployeeRepository(),
		),
	)
}
