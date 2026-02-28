package routes

import (
	"Apps-I_Desa_Backend/controllers"
	"Apps-I_Desa_Backend/middleware"
	"Apps-I_Desa_Backend/repositories"
	"Apps-I_Desa_Backend/services"
	"github.com/gofiber/fiber/v2"
)

func SetupSubDimensionRoutes(app *fiber.App) {
	subDimensionRepo := repositories.NewSubDimensionRepository()
	subDimensionService := services.NewSubDimensionService(subDimensionRepo)
	subDimensionController := controllers.NewSubDimensionController(subDimensionService)

	api := app.Group("/api/sub-dimensions")

	// Apply JWT middleware to all sub-dimension routes
	api.Use(middleware.JWTAuth())

	// Pendidikan routes
	api.Post("/pendidikan", subDimensionController.CreateSubDimensionPendidikan)

	// Kesehatan routes
	api.Post("/kesehatan", subDimensionController.CreateSubDimensionKesehatan)

	// Utilitas Dasar routes
	api.Post("/utilitas-dasar", subDimensionController.CreateSubDimensionUtilitasDasar)

	// Aktivitas routes
	api.Post("/aktivitas", subDimensionController.CreateSubDimensionAktivitas)

	// Fasilitas Masyarakat routes
	api.Post("/fasilitas-masyarakat", subDimensionController.CreateSubDimensionFasilitasMasyarakat)

	// Produksi Desa routes
	api.Post("/produksi-desa", subDimensionController.CreateSubDimensionProduksiDesa)

	// Fasilitas Pendukung Ekonomi routes
	api.Post(
		"/fasilitas-pendukung-ekonomi",
		subDimensionController.CreateSubDimensionFasilitasPendukungEkonomi,
	)

	// Pengelolaan Lingkungan routes
	api.Post(
		"/pengelolaan-lingkungan",
		subDimensionController.CreateSubDimensionPengelolaanLingkungan,
	)

	// Penanggulangan Bencana routes
	api.Post(
		"/penanggulangan-bencana",
		subDimensionController.CreateSubDimensionPenanggulanganBencana,
	)

	// Kondisi Akses Jalan routes
	api.Post("/kondisi-akses-jalan", subDimensionController.CreateSubDimensionKondisiAksesJalan)

	// Kemudahan Akses routes
	api.Post("/kemudahan-akses", subDimensionController.CreateSubDimensionKemudahanAkses)

	// Kelembagaan Pelayanan Desa routes
	api.Post(
		"/kelembagaan-pelayanan-desa",
		subDimensionController.CreateSubDimensionKelembagaanPelayananDesa,
	)

	// Tata Kelola Keuangan Desa routes
	api.Post(
		"/tata-kelola-keuangan-desa",
		subDimensionController.CreateSubDimensionTataKelolaKeuanganDesa,
	)
}
