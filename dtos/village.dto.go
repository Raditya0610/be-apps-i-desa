package dtos

import "github.com/google/uuid"

type AddVillageRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateVillageResponse struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Message string    `json:"message"`
}
