package models

import (
	"Zynto/internal/services/models/enums"
	"time"
)

type Service struct {
	ID            string                   `json:"id"`
	Name          string                   `json:"name"`
	Description   string                   `json:"description"`
	GenderService enums.GenderEnumServices `json:"gender_service"`
	Photo         string                   `json:"photo"`
	Active        bool                     `json:"active"`
	Price         int                      `json:"price"`
	Duration      int                      `json:"duration"`
	Category      string                   `json:"category"`
	CompanyId     string                   `json:"company_id"`
	EmployeeId    string                   `json:"employee_id"`
	CreatedAt     time.Time                `json:"created_at"`
	UpdatedAt     time.Time                `json:"updated_at"`
}
