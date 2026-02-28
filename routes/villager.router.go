package routes

import (
	"Apps-I_Desa_Backend/controllers"
	"Apps-I_Desa_Backend/middleware"
	"Apps-I_Desa_Backend/repositories"
	"Apps-I_Desa_Backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupVillagerRoutes(app *fiber.App) {
	villagerRepo := repositories.NewVillagerRepository()
	villagerService := services.NewVillagerService(villagerRepo)
	villagerController := controllers.NewVillagerController(villagerService)

	api := app.Group("/api/villagers")
	// Apply JWT middleware to all villager routes
	api.Use(middleware.JWTAuth())

	api.Post("/", villagerController.CreateVillager)
	api.Get("/", villagerController.GetAllVillagers)
	api.Get("/:nik", villagerController.GetVillager)
	api.Put("/:nik", villagerController.UpdateVillager)
	api.Delete("/:nik", villagerController.DeleteVillager)
}
