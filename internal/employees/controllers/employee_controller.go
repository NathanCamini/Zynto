package employees

import (
	"Zynto/internal/employees/models"
	"Zynto/internal/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type EmployeesController struct {
	employeeService models.EmployeeService
}

func NewEmployeesController(employeeService models.EmployeeService) *EmployeesController {
	return &EmployeesController{employeeService: employeeService}
}

func (e *EmployeesController) RegisterRoutes() http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/employees", func(r chi.Router) {
		r.Post("/", e.createEmployee())
		r.Get("/", e.getAllEmployees())
		r.Get("/{userID}", e.getEmployee())
		r.Put("/{userID}", e.updateEmployee())
		r.Delete("/{userID}", e.deleteUser())
	})

	return r
}

func (e *EmployeesController) createEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var employee models.Employee

		if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
			utils.SendJson(w, utils.Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		createdEmployee, err := e.employeeService.CreateEmployee(&employee)
		if err != nil {
			utils.SendJson(w, utils.Response{Error: "errrororr"}, http.StatusBadRequest)
			return
		}

		utils.SendJson(w, utils.Response{Data: createdEmployee}, http.StatusCreated)
	}
}

func (e *EmployeesController) getEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")

		employee, err := e.employeeService.GetEmployee(userID)
		if err != nil {
			utils.SendJson(w, utils.Response{Error: err.Error()}, http.StatusNotFound)
			return
		}
		utils.SendJson(w, utils.Response{Data: employee}, http.StatusOK)
	}
}

func (e *EmployeesController) getAllEmployees() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employees, err := e.employeeService.GetAllEmployees()

		if err != nil {
			utils.SendJson(w, utils.Response{Error: err.Error()}, http.StatusInternalServerError)
			return
		}

		utils.SendJson(w, utils.Response{Data: employees}, http.StatusOK)
	}
}

func (e *EmployeesController) updateEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		var employee models.Employee

		if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
			utils.SendJson(w, utils.Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if err := e.employeeService.UpdateEmployee(userID, &employee); err != nil {
			utils.SendJson(w, utils.Response{Error: err.Error()}, http.StatusNotFound)
			return
		}

		utils.SendJson(w, utils.Response{Data: employee}, http.StatusOK)
	}
}

func (e *EmployeesController) deleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")

		err := e.employeeService.DeleteEmployee(userID)
		if err != nil {
			utils.SendJson(w, utils.Response{Error: err.Error()}, http.StatusNotFound)
			return
		}

		utils.SendJson(w, utils.Response{}, http.StatusOK)
	}
}
