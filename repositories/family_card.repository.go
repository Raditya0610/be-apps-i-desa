package repositories

import (
	"Apps-I_Desa_Backend/config"
	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FamilyCardRepository struct {
	DB *gorm.DB
}

func NewFamilyCardRepository() *FamilyCardRepository {
	return &FamilyCardRepository{
		DB: config.DB,
	}
}

func (r *FamilyCardRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}

func (r *FamilyCardRepository) CreateWithTx(tx *gorm.DB, familyCard *models.FamilyCard) error {
	return tx.Create(familyCard).Error
}

func (r *FamilyCardRepository) GetNIKAndAddressByNIK(nik string) (*dtos.GetAllFamilyMember, error) {
	var familyCard models.FamilyCard
	err := r.DB.Where("nik = ?", nik).Select("nik", "alamat").First(&familyCard).Error
	if err != nil {
		return nil, err
	}

	return &dtos.GetAllFamilyMember{
		NIK:     familyCard.NIK,
		Address: familyCard.Alamat,
	}, nil
}

type FamilyCardSummary struct {
	NIK            string
	KepalaKeluarga string
	TotalMembers   int
}

func (r *FamilyCardRepository) GetAllFamilyCardsByVillageID(
	villageID *uuid.UUID,
) ([]*FamilyCardSummary, error) {
	var results []*FamilyCardSummary
	err := r.DB.Model(&models.FamilyCard{}).
		Select(`family_cards.nik,
			COALESCE(v_head.nama_lengkap, '') AS kepala_keluarga,
			COUNT(v_all.nik) AS total_members`).
		Joins("LEFT JOIN villagers v_head ON v_head.family_card_id = family_cards.nik AND v_head.status_hubungan = 'Kepala Keluarga'").
		Joins("LEFT JOIN villagers v_all ON v_all.family_card_id = family_cards.nik").
		Where("family_cards.village_id = ?", villageID).
		Group("family_cards.nik, v_head.nama_lengkap").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *FamilyCardRepository) GetFamilyCardByNIK(nik *string) (*models.FamilyCard, error) {
	var familyCard models.FamilyCard
	err := r.DB.Where("nik = ?", nik).First(&familyCard).Error
	if err != nil {
		return nil, err
	}
	return &familyCard, nil
}

func (r *FamilyCardRepository) CountAllFamilyCardByVillageID(villageID *uuid.UUID) (int64, error) {
	var count int64
	err := r.DB.Model(&models.FamilyCard{}).Where("village_id = ?", villageID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *FamilyCardRepository) CountDistinctRT(villageID *uuid.UUID) (int64, error) {
	var count int64
	err := r.DB.Model(&models.FamilyCard{}).
		Where("village_id = ?", villageID).
		Distinct("rt").
		Count(&count).
		Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *FamilyCardRepository) CountDistinctRW(villageID *uuid.UUID) (int64, error) {
	var count int64
	err := r.DB.Model(&models.FamilyCard{}).
		Where("village_id = ?", villageID).
		Distinct("rw").
		Count(&count).
		Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *FamilyCardRepository) CountDistinctKelurahan(villageID *uuid.UUID) (int64, error) {
	var count int64
	err := r.DB.Model(&models.FamilyCard{}).
		Where("village_id = ?", villageID).
		Distinct("kelurahan").
		Count(&count).
		Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *FamilyCardRepository) CountDistinctKecamatan(villageID *uuid.UUID) (int64, error) {
	var count int64
	err := r.DB.Model(&models.FamilyCard{}).
		Where("village_id = ?", villageID).
		Distinct("kecamatan").
		Count(&count).
		Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
