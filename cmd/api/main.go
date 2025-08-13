package main

import (
	di "Zynto/internal/employees/DI"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}
	slog.Info("all systems offline")
}

func run() error {

	employeesController := di.NewAppEmployeesController()

	handler := employeesController.RegisterRoutes()

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
