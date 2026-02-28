package controllers

import (
	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type VillageController struct {
	villageService *services.VillageService
	validate       *validator.Validate
}

func NewVillageController(villageService *services.VillageService) *VillageController {
	return &VillageController{
		villageService: villageService,
		validate:       validator.New(),
	}
}

func (c *VillageController) CreateVillage(ctx *fiber.Ctx) error {
	var request dtos.AddVillageRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := c.validate.Struct(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.villageService.CreateVillage(&request)
	if err != nil {
		if err.Error() == "failed to create village" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create village",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to commit transaction" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to commit transaction",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create village",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}
