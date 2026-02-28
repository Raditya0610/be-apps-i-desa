package routes

import (
	"Apps-I_Desa_Backend/controllers"
	"Apps-I_Desa_Backend/middleware"
	"Apps-I_Desa_Backend/repositories"
	"Apps-I_Desa_Backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupDashboardRoutes(app *fiber.App) {
	villagerRepo := repositories.NewVillagerRepository()
	familyCardRepo := repositories.NewFamilyCardRepository()
	dashboardService := services.NewDashboardService(villagerRepo, familyCardRepo)
	dashboardController := controllers.NewDashboardController(dashboardService)

	api := app.Group("/api/dashboard")
	// Apply JWT middleware to all dashboard routes
	api.Use(middleware.JWTAuth())

	api.Get("/", dashboardController.GetDashboardData)
}
