package models

import (
	"time"

	"github.com/google/uuid"
)

type Villager struct {
	NIK              string    `gorm:"primaryKey;size:16"`
	NamaLengkap      string    `gorm:"size:100;not null"`
	JenisKelamin     string    `gorm:"size:10;not null"` // e.g., "Laki-laki", "Perempuan"
	TempatLahir      string    `gorm:"size:100;not null"`
	TanggalLahir     time.Time `gorm:"not null"`         // Use time.Time for date fields
	Agama            string    `gorm:"size:20;not null"` // e.g., "Islam", "Kristen", etc.
	Pendidikan       string    `gorm:"size:50;not null"` // e.g., "SD", "SMP", "SMA", "Sarjana"
	Pekerjaan        string    `gorm:"size:50;not null"` // e.g., "Petani", "Guru", "Dokter"
	StatusPerkawinan string    `gorm:"size:20;not null"` // e.g., "
	StatusHubungan   string    `gorm:"size:20;not null"` // e.g., "Kepala Keluarga", "Ang
	Kewarganegaraan  string    `gorm:"size:20;not null"` // e.g., "WNI", "WNA"
	NomorPaspor      *string   `gorm:"size:20"`          // Optional field for foreign nationals
	NomorKitas       *string   `gorm:"size:20"`          // Optional field for foreign nationals
	NamaAyah         string    `gorm:"size:100;not null"`
	NamaIbu          string    `gorm:"size:100;not null"`
	VillageID        uuid.UUID `gorm:"type:uuid;not null"`
	FamilyCardID     string    `gorm:"not null"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`

	// Belongs to relationships
	Village    Village    `gorm:"foreignKey:VillageID"`
	FamilyCard FamilyCard `gorm:"foreignKey:FamilyCardID;references:NIK"`
}
