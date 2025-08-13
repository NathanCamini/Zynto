package models

import (
	enums "Zynto/internal/employees/models/enums"
	"time"
)

type Employee struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Document    string           `json:"document"`
	Gender      enums.GenderEnum `json:"gender"`
	PhoneNumber string           `json:"phone_number"`
	Photo       string           `json:"photo"`
	Active      bool             `json:"active"`
	CompanyId   string           `json:"company_id"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}
