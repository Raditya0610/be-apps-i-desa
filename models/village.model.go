package models

import "github.com/google/uuid"

type Village struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string    `gorm:"size:100;not null"`

	// One-to-many relationships
	Users                               []User                                `gorm:"foreignKey:VillageID"`
	Villagers                           []Villager                            `gorm:"foreignKey:VillageID"`
	FamilyCards                         []FamilyCard                          `gorm:"foreignKey:VillageID"`
	SubDimensiPendidikan                []SubDimensiPendidikan                `gorm:"foreignKey:VillageID"`
	SubDimensiKesehatan                 []SubDimensiKesehatan                 `gorm:"foreignKey:VillageID"`
	SubDimensiUtilitasDasar             []SubDimensiUtilitasDasar             `gorm:"foreignKey:VillageID"`
	SubDimensiAktivitas                 []SubDimensiAktivitas                 `gorm:"foreignKey:VillageID"`
	SubDimensiFasilitasMasyarakat       []SubDimensiFasilitasMasyarakat       `gorm:"foreignKey:VillageID"`
	SubDimensiProduksiDesa              []SubDimensiProduksiDesa              `gorm:"foreignKey:VillageID"`
	SubDimensiFasilitasPendukungEkonomi []SubDimensiFasilitasPendukungEkonomi `gorm:"foreignKey:VillageID"`
	SubDimensiPengelolaanLingkungan     []SubDimensiPengelolaanLingkungan     `gorm:"foreignKey:VillageID"`
	SubDimensiPenanggulanganBencana     []SubDimensiPenanggulanganBencana     `gorm:"foreignKey:VillageID"`
	SubDimensiKondisiAksesJalan         []SubDimensiKondisiAksesJalan         `gorm:"foreignKey:VillageID"`
	SubDimensiKemudahanAkses            []SubDimensiKemudahanAkses            `gorm:"foreignKey:VillageID"`
	SubDimensiKelembagaanPelayananDesa  []SubDimensiKelembagaanPelayananDesa  `gorm:"foreignKey:VillageID"`
	SubDimensiTataKelolaKeuanganDesa    []SubDimensiTataKelolaKeuanganDesa    `gorm:"foreignKey:VillageID"`
}
