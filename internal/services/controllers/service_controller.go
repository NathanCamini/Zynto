package controllers

import (
	"Zynto/internal/services/models"
	"Zynto/internal/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type ServicesController struct {
	serviceService models.ServiceService
}

func NewServicesController(serviceService models.ServiceService) *ServicesController {
	return &ServicesController{serviceService: serviceService}
}

func (s *ServicesController) RegisterRoutes(r chi.Router) {
	r.Route("/services", func(r chi.Router) {
		r.Post("/", s.createService())
		r.Get("/", s.getAllServices())
		r.Get("/{serviceID}", s.getServiceByID())
		r.Put("/{serviceID}", s.updateService())
		r.Delete("/{serviceID}", s.deleteService())
	})
}

func (s *ServicesController) createService() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var service models.Service

		if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
			utils.SendJson(w, utils.Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		createdService, err := s.serviceService.CreateService(&service)
		if err != nil {
			utils.SendJson(w, utils.Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		utils.SendJson(w, utils.Response{Data: createdService}, http.StatusCreated)
	}
}

func (s *ServicesController) getServiceByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serviceID := chi.URLParam(r, "serviceID")

		service, err := s.serviceService.GetServiceByID(serviceID)
		if err != nil {
			utils.SendJson(w, utils.Response{Error: err.Error()}, http.StatusNotFound)
			return
		}

		utils.SendJson(w, utils.Response{Data: service}, http.StatusOK)
	}
}

func (s *ServicesController) getAllServices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		services, err := s.serviceService.GetAllServices()
		if err != nil {
			utils.SendJson(w, utils.Response{Error: err.Error()}, http.StatusInternalServerError)
			return
		}

		utils.SendJson(w, utils.Response{Data: services}, http.StatusOK)
	}
}

func (s *ServicesController) updateService() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serviceID := chi.URLParam(r, "serviceID")
		var service models.Service

		if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
			utils.SendJson(w, utils.Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if err := s.serviceService.UpdateService(serviceID, &service); err != nil {
			utils.SendJson(w, utils.Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		utils.SendJson(w, utils.Response{Data: "service updated successfully"}, http.StatusOK)
	}
}

func (s *ServicesController) deleteService() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serviceID := chi.URLParam(r, "serviceID")

		if err := s.serviceService.DeleteService(serviceID); err != nil {
			utils.SendJson(w, utils.Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		utils.SendJson(w, utils.Response{Data: "service deleted successfully"}, http.StatusOK)
	}
}
