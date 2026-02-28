package dtos

import "github.com/google/uuid"

type RegisterRequest struct {
	Username  string    `json:"username"   validate:"required"`
	Password  string    `json:"password"   validate:"required"`
	VillageID uuid.UUID `json:"village_id" validate:"required,uuid"`
}
