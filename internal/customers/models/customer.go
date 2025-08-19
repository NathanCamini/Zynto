package models

import (
	"Zynto/internal/customers/models/enums"
	"time"
)

type Customer struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Gender    enums.GenderEnum `json:"gender"`
	Phone     string           `json:"phone"`
	Email     string           `json:"email"`
	CompanyId string           `json:"company_id"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
