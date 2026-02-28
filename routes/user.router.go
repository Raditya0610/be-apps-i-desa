package routes

import (
	"Apps-I_Desa_Backend/controllers"
	"Apps-I_Desa_Backend/repositories"
	"Apps-I_Desa_Backend/services"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	villageRepo := repositories.NewVillageRepository()
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo, villageRepo)
	userController := controllers.NewUserController(userService)

	userRoutes := app.Group("/api/users")
	userRoutes.Post("/register", userController.Register)
}
