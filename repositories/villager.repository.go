package repositories

import (
	"time"

	"Apps-I_Desa_Backend/config"
	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VillagerRepository struct {
	DB *gorm.DB
}

func NewVillagerRepository() *VillagerRepository {
	return &VillagerRepository{
		DB: config.DB,
	}
}

func (r *VillagerRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}

func (r *VillagerRepository) FindVillagerByNIK(nik *string) (*models.Villager, error) {
	var villager models.Villager
	err := r.DB.Where("nik = ?", &nik).First(&villager).Error
	if err != nil {
		return nil, err
	}
	return &villager, nil
}

func (r *VillagerRepository) CreateVillagerWithTx(tx *gorm.DB, villager *models.Villager) error {
	return tx.Create(villager).Error
}

func (r *VillagerRepository) GetVillagersByFamilyCardNIK(
	familyCardNIK *string,
) ([]*dtos.GetFamilyMember, error) {
	var villagers []*models.Villager
	err := r.DB.Where("family_card_id = ?", &familyCardNIK).Find(&villagers).Error
	if err != nil {
		return nil, err
	}

	var familyMembers []*dtos.GetFamilyMember
	for _, villager := range villagers {
		age := calculateAge(villager.TanggalLahir)
		familyMember := &dtos.GetFamilyMember{
			NIK:            villager.NIK,
			Name:           villager.NamaLengkap,
			StatusHubungan: villager.StatusHubungan,
			Age:            age,
			JenisKelamin:   villager.JenisKelamin,
			Pendidikan:     villager.Pendidikan,
			Pekerjaan:      villager.Pekerjaan,
		}
		familyMembers = append(familyMembers, familyMember)
	}

	return familyMembers, nil
}

func (r *VillagerRepository) UpdateVillagerWithTx(tx *gorm.DB, villager *models.Villager) error {
	return tx.Save(villager).Error
}

func (r *VillagerRepository) DeleteVillagerWithTx(tx *gorm.DB, nik *string) error {
	return tx.Delete(&models.Villager{}, "nik = ?", *nik).Error
}

func (r *VillagerRepository) CountAllVillagerByVillageID(villageID *uuid.UUID) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Villager{}).Where("village_id = ?", villageID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *VillagerRepository) CountAllLakiLakiVillager(villageID *uuid.UUID) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Villager{}).
		Where("village_id = ? AND jenis_kelamin = ?", villageID, "Laki-laki").
		Count(&count).
		Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *VillagerRepository) GetAverageAge(villageID *uuid.UUID) (float32, error) {
	var averageAge float32
	err := r.DB.Model(&models.Villager{}).
		Where("village_id = ?", villageID).
		Select("AVG(EXTRACT(YEAR FROM AGE(CURRENT_DATE, tanggal_lahir)))").
		Scan(&averageAge).Error
	if err != nil {
		return 0, err
	}
	return averageAge, nil
}

func (r *VillagerRepository) CountAllKepalaKeluarga(villageID *uuid.UUID) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Villager{}).
		Where("village_id = ? AND status_hubungan = ?", villageID, "Kepala Keluarga").
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *VillagerRepository) GetAllVillagers(villageID *uuid.UUID) ([]*dtos.GetAllVillagersResponse, error) {
	var villagers []*dtos.GetAllVillagersResponse
	err := r.DB.Model(&models.Villager{}).
		Where("village_id = ?", villageID).
		Select("nik", "nama_lengkap").
		Scan(&villagers).Error
	if err != nil {
		return nil, err
	}
	return villagers, nil
}

func calculateAge(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()

	// Adjust if birthday hasn't occurred this year
	if now.YearDay() < birthDate.YearDay() {
		age--
	}

	return age
}
