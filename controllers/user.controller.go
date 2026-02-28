package controllers

import (
	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *services.UserService
	validate    *validator.Validate
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
		validate:    validator.New(),
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var request dtos.RegisterRequest

	// Parse request body
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validate request data
	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	// Process registration
	response, err := c.userService.Register(&request)
	if err != nil {
		if err.Error() == "username already registered" {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "Username already registered",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to find existing user" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to find existing user",
				"error":   err.Error(),
			})
		} else if err.Error() == "error hashing password" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error hashing password",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to create user" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create user",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to commit transaction" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to commit transaction",
				"error":   err.Error(),
			})
		} else if err.Error() == "village not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Village not found",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}
