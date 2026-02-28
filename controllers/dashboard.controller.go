package controllers

import (
	"Apps-I_Desa_Backend/services"
	"github.com/gofiber/fiber/v2"
)

type DashboardController struct {
	dashboardService *services.DashboardService
}

func NewDashboardController(userService *services.DashboardService) *DashboardController {
	return &DashboardController{
		dashboardService: userService,
	}
}

func (c *DashboardController) GetDashboardData(ctx *fiber.Ctx) error {
	response, err := c.dashboardService.GetDashboardData(ctx)
	if err != nil {
		if err.Error() == "village ID is empty" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Village ID is empty",
				"error":   err.Error(),
			})
		} else if err.Error() == "invalid village ID format" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid village ID format",
				"error":   err.Error(),
			})
		} else if err.Error() == "error counting family cards" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to count family cards",
				"error":   err.Error(),
			})
		} else if err.Error() == "error counting RT" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to count RT",
				"error":   err.Error(),
			})
		} else if err.Error() == "error counting RW" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to count RW",
				"error":   err.Error(),
			})
		} else if err.Error() == "error counting kelurahan" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to count kelurahan",
				"error":   err.Error(),
			})
		} else if err.Error() == "error counting kecamatan" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to count kecamatan",
				"error":   err.Error(),
			})
		} else if err.Error() == "error counting villagers" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to count villagers",
				"error":   err.Error(),
			})
		} else if err.Error() == "error counting male villagers" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to count male villagers",
				"error":   err.Error(),
			})
		} else if err.Error() == "error getting average age" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to get average age",
				"error":   err.Error(),
			})
		} else if err.Error() == "error counting kepala keluarga" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to count kepala keluarga",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get dashboard data",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
