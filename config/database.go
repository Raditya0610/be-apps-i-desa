package config

import (
	"fmt"
	"os"
	"time"

	models2 "Apps-I_Desa_Backend/models"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DB is a global variable that holds the database connection
var DB *gorm.DB

// ConnectDB initializes the database connection and returns it
func ConnectDB() *gorm.DB {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Warn("Error loading .env file, using environment variables")
	}

	// Validate required environment variables
	// requiredEnvVars := []string{"DB_HOST", "DB_USERNAME", "DB_PASSWORD", "DB_NAME", "DB_PORT"}
	// for _, envVar := range requiredEnvVars {
	// 	if os.Getenv(envVar) == "" {
	// 		log.Fatalf("Required environment variable %s is not set", envVar)
	// 	}
	// }

	// Set default values for optional environment variables
	sslMode := os.Getenv("DB_SSL")
	if sslMode == "" {
		sslMode = "disable"
	}

	timeZone := os.Getenv("DB_TIMEZONE")
	if timeZone == "" {
		timeZone = "UTC"
	}

	// Build connection string — prefer DATABASE_URL (full Service URI) if set
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf(
			"host=%s user=%s dbname=%s port=%s sslmode=%s TimeZone=%s password=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
			sslMode,
			timeZone,
			os.Getenv("DB_PASSWORD"),
		)
	}

	// Connect to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: false,
		},
	})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database connection pool: ", err)
	}

	// Configure connection pool for serverless/limited connections
	// In serverless environments, each instance should use minimal connections
	sqlDB.SetMaxOpenConns(2)                   // Maximum 2 connections per instance
	sqlDB.SetMaxIdleConns(1)                   // Keep 1 idle connection for reuse
	sqlDB.SetConnMaxLifetime(10 * time.Minute) // Shorter lifetime for serverless
	sqlDB.SetConnMaxIdleTime(3 * time.Minute)  // Shorter idle time to free resources

	// Run migrations (Skip in production to avoid cold-start lag and heavy Aiven DB queries)
	env := os.Getenv("APP_ENV")
	if env != "production" {
		log.Info("Running Database Migrations...")
		if err := migrateDB(DB); err != nil {
			log.Fatal("Database migration failed: ", err)
		}
	} else {
		log.Info("Skipping AutoMigrate for Production Environment")
	}

	log.Info("Successfully connected to the database", dsn)
	return DB
}

func migrateDB(db *gorm.DB) error {
	m := []interface{}{
		&models2.User{},
		&models2.Village{},
		&models2.Villager{},
		&models2.FamilyCard{},
		&models2.SubDimensiAktivitas{},
		&models2.SubDimensiFasilitasMasyarakat{},
		&models2.SubDimensiFasilitasPendukungEkonomi{},
		&models2.SubDimensiKelembagaanPelayananDesa{},
		&models2.SubDimensiKemudahanAkses{},
		&models2.SubDimensiKondisiAksesJalan{},
		&models2.SubDimensiPendidikan{},
		&models2.SubDimensiKesehatan{},
		&models2.SubDimensiUtilitasDasar{},
		&models2.SubDimensiProduksiDesa{},
		&models2.SubDimensiPengelolaanLingkungan{},
		&models2.SubDimensiPenanggulanganBencana{},
		&models2.SubDimensiTataKelolaKeuanganDesa{},
	}

	return db.AutoMigrate(m...)
}

func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Error("Failed to get database connection: ", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Error("Error closing database connection: ", err)
		} else {
			log.Info("Database connection closed")
		}
	}
}
