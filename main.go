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

			// Default ke 500
			code := fiber.StatusInternalServerError
			errorMessage := "Internal Server Error"

			// Cek apakah ini error internal bawaan Fiber (seperti 404 Not Found)
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				errorMessage = e.Message
			} else if env == "development" {
				// Kalau bukan error HTTP biasa dan env=development, tampilkan error asli golang
				errorMessage = err.Error()
			}

			// Custom 404 Message buat DT-07
			if code == fiber.StatusNotFound {
				errorMessage = "Endpoint tidak ditemukan (404 Not Found)"
			}

			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": errorMessage,
			})
		},
	})

	app.Use(logger.New())
	app.Use(recover.New())

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
