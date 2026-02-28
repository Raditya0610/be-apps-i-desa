package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	models "Apps-I_Desa_Backend/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env not found, using system environment variables")
	}

	db := connectDB()
	log.Println("✅ Connected to DB")

	// Resolve CSV directory relative to project root
	baseDir := "../DB"
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		baseDir = "DB"
	}
	log.Printf("📂 Using CSV directory: %s\n", baseDir)

	seedVillages(db, baseDir+"/villages.csv")
	seedFamilyCards(db, baseDir+"/family_cards.csv")
	seedUsers(db, baseDir+"/users.csv")
	seedVillagers(db, baseDir+"/villagers.csv")

	log.Println("🎉 Seeding complete!")
}

func connectDB() *gorm.DB {
	// Prefer DATABASE_URL if set
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		sslMode := os.Getenv("DB_SSL")
		if sslMode == "" {
			sslMode = "disable"
		}
		timezone := os.Getenv("DB_TIMEZONE")
		if timezone == "" {
			timezone = "UTC"
		}
		dsn = fmt.Sprintf(
			"host=%s user=%s dbname=%s port=%s sslmode=%s TimeZone=%s password=%s",
			os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"), sslMode, timezone, os.Getenv("DB_PASSWORD"),
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: false},
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}
	return db
}

// villages.csv: id,name
func seedVillages(db *gorm.DB, path string) {
	rows := readCSV(path)
	var records []models.Village
	for _, row := range rows {
		if len(row) < 2 {
			continue
		}
		id, err := uuid.Parse(strings.TrimSpace(row[0]))
		if err != nil {
			continue
		}
		records = append(records, models.Village{ID: id, Name: strings.TrimSpace(row[1])})
	}
	res := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&records)
	log.Printf("🏘️  Villages: %d inserted, error: %v", res.RowsAffected, res.Error)
}

// family_cards.csv: NIK,Alamat,RT,RW,Kelurahan,Kecamatan,KabupatenKota,KodePos,Provinsi,CreatedAt,UpdatedAt,VillageID
func seedFamilyCards(db *gorm.DB, path string) {
	rows := readCSV(path)
	var records []models.FamilyCard
	for _, row := range rows {
		if len(row) < 12 {
			continue
		}
		vid, err := uuid.Parse(strings.TrimSpace(row[11]))
		if err != nil {
			continue
		}
		records = append(records, models.FamilyCard{
			NIK: strings.TrimSpace(row[0]), Alamat: strings.TrimSpace(row[1]),
			RT: strings.TrimSpace(row[2]), RW: strings.TrimSpace(row[3]),
			Kelurahan: strings.TrimSpace(row[4]), Kecamatan: strings.TrimSpace(row[5]),
			KabupatenKota: strings.TrimSpace(row[6]), KodePos: strings.TrimSpace(row[7]),
			Provinsi: strings.TrimSpace(row[8]), VillageID: vid,
		})
	}
	res := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&records)
	log.Printf("📋 Family Cards: %d inserted, error: %v", res.RowsAffected, res.Error)
}

// users.csv: username,password,village_id
func seedUsers(db *gorm.DB, path string) {
	rows := readCSV(path)
	var records []models.User
	for _, row := range rows {
		if len(row) < 3 {
			continue
		}
		vid, err := uuid.Parse(strings.TrimSpace(row[2]))
		if err != nil {
			continue
		}
		records = append(records, models.User{
			Username: strings.TrimSpace(row[0]), Password: strings.TrimSpace(row[1]), VillageID: vid,
		})
	}
	res := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&records)
	log.Printf("👤 Users: %d inserted, error: %v", res.RowsAffected, res.Error)
}

// villagers.csv: NIK,NamaLengkap,JenisKelamin,TempatLahir,TanggalLahir,Agama,Pendidikan,
//
//	Pekerjaan,StatusPerkawinan,StatusHubungan,Kewarganegaraan,NomorPaspor,NomorKitas,
//	NamaAyah,NamaIbu,VillageID,FamilyCardID,CreatedAt,UpdatedAt
func seedVillagers(db *gorm.DB, path string) {
	rows := readCSV(path)
	var records []models.Villager
	for _, row := range rows {
		if len(row) < 17 {
			continue
		}
		tgl, err := parseTime(strings.TrimSpace(row[4]))
		if err != nil {
			log.Printf("⚠️  Skip villager %q bad date %q", row[0], row[4])
			continue
		}
		vid, err := uuid.Parse(strings.TrimSpace(row[15]))
		if err != nil {
			continue
		}
		var paspor, kitas *string
		if v := strings.TrimSpace(row[11]); v != "" && v != "-" {
			paspor = &v
		}
		if v := strings.TrimSpace(row[12]); v != "" && v != "-" {
			kitas = &v
		}
		records = append(records, models.Villager{
			NIK: strings.TrimSpace(row[0]), NamaLengkap: strings.TrimSpace(row[1]),
			JenisKelamin: strings.TrimSpace(row[2]), TempatLahir: strings.TrimSpace(row[3]),
			TanggalLahir: tgl, Agama: strings.TrimSpace(row[5]),
			Pendidikan: strings.TrimSpace(row[6]), Pekerjaan: strings.TrimSpace(row[7]),
			StatusPerkawinan: strings.TrimSpace(row[8]), StatusHubungan: strings.TrimSpace(row[9]),
			Kewarganegaraan: strings.TrimSpace(row[10]), NomorPaspor: paspor, NomorKitas: kitas,
			NamaAyah: strings.TrimSpace(row[13]), NamaIbu: strings.TrimSpace(row[14]),
			VillageID: vid, FamilyCardID: strings.TrimSpace(row[16]),
		})
	}
	total := int64(0)
	for i := 0; i < len(records); i += 50 {
		end := i + 50
		if end > len(records) {
			end = len(records)
		}
		batch := records[i:end]
		res := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&batch)
		if res.Error != nil {
			log.Printf("⚠️  Batch %d-%d error: %v", i, end, res.Error)
		}
		total += res.RowsAffected
	}
	log.Printf("👥 Villagers: %d inserted out of %d total", total, len(records))
}

func readCSV(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("❌ Cannot open %s: %v", path, err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalf("❌ Cannot parse CSV %s: %v", path, err)
	}
	var clean [][]string
	for _, row := range rows {
		if strings.TrimSpace(strings.Join(row, "")) != "" {
			clean = append(clean, row)
		}
	}
	return clean
}

var timeFormats = []string{
	"2006-01-02 15:04:05.999999 -07:00",
	"2006-01-02 15:04:05.999999 +00:00",
	"2006-01-02 15:04:05 +00:00",
	"2006-01-02T15:04:05Z",
	"2006-01-02 15:04:05",
	"2006-01-02",
}

func parseTime(s string) (time.Time, error) {
	for _, layout := range timeFormats {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("cannot parse %q", s)
}
