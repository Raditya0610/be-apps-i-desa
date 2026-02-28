package repositories

import (
	"Apps-I_Desa_Backend/config"
	"Apps-I_Desa_Backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VillageRepository struct {
	DB *gorm.DB
}

func NewVillageRepository() *VillageRepository {
	return &VillageRepository{
		DB: config.DB,
	}
}

func (r *VillageRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}

func (r *VillageRepository) CreateVillageWithTx(tx *gorm.DB, village *models.Village) error {
	return tx.Create(village).Error
}

func (r *VillageRepository) FindVillageByID(uuid *uuid.UUID) error {
	var village models.Village
	if err := r.DB.First(&village, "id = ?", uuid).Error; err != nil {
		return err
	}
	return nil
}

func (r *VillageRepository) FindVillageByName(name string) *models.Village {
	var village models.Village
	if err := r.DB.First(&village, "name = ?", name).Error; err != nil {
		return nil
	}
	return &village
}
