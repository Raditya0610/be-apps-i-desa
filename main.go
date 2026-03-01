package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"Apps-I_Desa_Backend/config"
	"Apps-I_Desa_Backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func setupRoutes(app *fiber.App) {
	routes.SetupUserRoutes(app)
	routes.SetupAuthRoutes(app)
	routes.SetupSubDimensionRoutes(app)
	routes.SetupVillagerRoutes(app)
	routes.SetupVillageRoutes(app)
	routes.SetupFamilyCardRoutes(app)
	routes.SetupDashboardRoutes(app)
}

func main() {
	config.ConnectDB()
	defer config.CloseDB()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// DT-03 Security Gate: Hide error details in production
			env := os.Getenv("APP_ENV")
			errorMessage := "Internal Server Error"

			if env == "development" {
				errorMessage = err.Error() // Show raw error only in dev
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": errorMessage,
			})
		},
	})

	app.Use(logger.New())
	app.Use(recover.New())

	// Custom 404 handler (untuk DT-07 Error Page Check)
	app.Use(func(c *fiber.Ctx) error {
		// Jika route tidak ditemukan dan bukan endpoint api yang valid
		if err := c.Next(); err != nil {
			return err
		}

		// Tangkap 404
		if c.Response().StatusCode() == 404 {
			return c.Status(404).JSON(fiber.Map{
				"success": false,
				"message": "Endpoint tidak ditemukan (404 Not Found)",
			})
		}
		return nil
	})

	// Read allowed origins from env (comma-separated), fallback to wildcard for local dev
	allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "*"
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,Cookie",
		AllowCredentials: true,
		ExposeHeaders:    "Set-Cookie",
	}))

	setupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Fatal("Error starting server: ", err)
		}
	}()

	log.Printf("Server started on port %s", port)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Apps-I Desa API!")
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatal("Server shutdown failed: ", err)
	}
}
