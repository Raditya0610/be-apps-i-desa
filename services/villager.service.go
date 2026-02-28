package services

import (
	"errors"
	"fmt"
	"log"
	"time"

	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/models"
	"Apps-I_Desa_Backend/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VillagerService struct {
	villagerRepo *repositories.VillagerRepository
}

func NewVillagerService(villagerRepo *repositories.VillagerRepository) *VillagerService {
	return &VillagerService{
		villagerRepo: villagerRepo,
	}
}

func (s *VillagerService) CreateVillager(
	request *dtos.AddVillagerRequest,
	ctx *fiber.Ctx,
) (*dtos.MessageResponse, error) {
	tx := s.villagerRepo.BeginTransaction()
	defer tx.Rollback()

	// Parse VillageID from string to UUID
	villageIDStr := fmt.Sprintf("%v", ctx.Locals("village"))
	if villageIDStr == "" {
		log.Println("Village ID is empty")
		return nil, errors.New("village ID is required")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Println("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID format")
	}

	// Parse TanggalLahir from string to time.Time
	tanggalLahir, err := time.Parse("2006-01-02", request.TanggalLahir)
	if err != nil {
		log.Println("Error parsing date:", err)
		return nil, errors.New("invalid date format, expected YYYY-MM-DD")
	}

	// Check if villager with the same NIK already exists
	existingVillager, err := s.villagerRepo.FindVillagerByNIK(&request.NIK)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Error checking existing villager:", err)
		return nil, errors.New("failed to check existing villager")
	}

	if existingVillager != nil {
		log.Println("Villager with the same NIK already exists")
		return nil, errors.New("villager with the same NIK already exists")
	}

	villager := &models.Villager{
		NIK:              request.NIK,
		NamaLengkap:      request.NamaLengkap,
		JenisKelamin:     request.JenisKelamin,
		TempatLahir:      request.TempatLahir,
		TanggalLahir:     tanggalLahir,
		Agama:            request.Agama,
		Pendidikan:       request.Pendidikan,
		Pekerjaan:        request.Pekerjaan,
		StatusPerkawinan: request.StatusPerkawinan,
		StatusHubungan:   request.StatusHubungan,
		Kewarganegaraan:  request.Kewarganegaraan,
		NamaAyah:         request.NamaAyah,
		NamaIbu:          request.NamaIbu,
		VillageID:        villageID,
		FamilyCardID:     request.FamilyCardID,
	}

	if request.NomorKitas != nil {
		villager.NomorKitas = request.NomorKitas
	}

	if request.NomorPaspor != nil {
		villager.NomorPaspor = request.NomorPaspor
	}

	if err := s.villagerRepo.CreateVillagerWithTx(tx, villager); err != nil {
		log.Println("Error creating villager:", err)
		return nil, errors.New("failed to create villager")
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}

	return &dtos.MessageResponse{
		Message: "Villager created successfully",
	}, nil
}

func (s *VillagerService) GetVillagerByNIK(nik *string) (*dtos.GetVillagerResponse, error) {
	villager, err := s.villagerRepo.FindVillagerByNIK(nik)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Villager not found:", err)
			return nil, errors.New("villager not found")
		}
		log.Println("Error finding villager:", err)
		return nil, errors.New("failed to find villager")
	}
	response := &dtos.GetVillagerResponse{
		NIK:              villager.NIK,
		NamaLengkap:      villager.NamaLengkap,
		JenisKelamin:     villager.JenisKelamin,
		TempatLahir:      villager.TempatLahir,
		TanggalLahir:     villager.TanggalLahir.Format("2006-01-02"),
		Agama:            villager.Agama,
		Pendidikan:       villager.Pendidikan,
		Pekerjaan:        villager.Pekerjaan,
		StatusPerkawinan: villager.StatusPerkawinan,
		StatusHubungan:   villager.StatusHubungan,
		Kewarganegaraan:  villager.Kewarganegaraan,
		NomorPaspor:      villager.NomorPaspor,
		NomorKitas:       villager.NomorKitas,
		NamaAyah:         villager.NamaAyah,
		NamaIbu:          villager.NamaIbu,
		FamilyCardID:     villager.FamilyCardID,
		VillageID:        villager.VillageID.String(),
	}
	return response, nil
}

func (s *VillagerService) GetAllVillagers(ctx *fiber.Ctx) ([]*dtos.GetAllVillagersResponse, error) {
	villageIDStr := fmt.Sprintf("%v", ctx.Locals("village"))
	if villageIDStr == "" || villageIDStr == "<nil>" {
		return nil, errors.New("village ID is required")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Println("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID format")
	}

	villagers, err := s.villagerRepo.GetAllVillagers(&villageID)
	if err != nil {
		log.Println("Error getting all villagers:", err)
		return nil, errors.New("failed to get all villagers")
	}

	return villagers, nil
}

func (s *VillagerService) UpdateVillager(
	nik *string,
	request *dtos.UpdateVillagerRequest,
) (*dtos.MessageResponse, error) {
	tx := s.villagerRepo.BeginTransaction()
	defer tx.Rollback()

	villager, err := s.villagerRepo.FindVillagerByNIK(nik)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Villager not found:", err)
			return nil, errors.New("villager not found")
		}
		log.Println("Error finding villager:", err)
		return nil, errors.New("failed to find villager")
	}

	if request.NIK != nil {
		// Check if villager with the same NIK already exists
		existingVillager, err := s.villagerRepo.FindVillagerByNIK(request.NIK)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Error checking existing villager:", err)
			return nil, errors.New("failed to check existing villager")
		}

		if existingVillager != nil && existingVillager.NIK != *nik {
			log.Println("Villager with the same NIK already exists")
			return nil, errors.New("villager with the same NIK already exists")
		}
		villager.NIK = *request.NIK
	}

	if request.NamaLengkap != nil {
		villager.NamaLengkap = *request.NamaLengkap
	}
	if request.JenisKelamin != nil {
		villager.JenisKelamin = *request.JenisKelamin
	}
	if request.TempatLahir != nil {
		villager.TempatLahir = *request.TempatLahir
	}
	if request.TanggalLahir != nil {
		// Validate TanggalLahir format
		if _, err := time.Parse("2006-01-02", *request.TanggalLahir); err != nil {
			log.Println("Error parsing date:", err)
			return nil, errors.New("invalid date format, expected YYYY-MM-DD")
		}
	}
	if request.Agama != nil {
		villager.Agama = *request.Agama
	}
	if request.Pendidikan != nil {
		villager.Pendidikan = *request.Pendidikan
	}
	if request.Pekerjaan != nil {
		villager.Pekerjaan = *request.Pekerjaan
	}
	if request.StatusPerkawinan != nil {
		villager.StatusPerkawinan = *request.StatusPerkawinan
	}
	if request.StatusHubungan != nil {
		villager.StatusHubungan = *request.StatusHubungan
	}
	if request.Kewarganegaraan != nil {
		villager.Kewarganegaraan = *request.Kewarganegaraan
	}
	if request.NamaAyah != nil {
		villager.NamaAyah = *request.NamaAyah
	}
	if request.NamaIbu != nil {
		villager.NamaIbu = *request.NamaIbu
	}
	if request.NomorKitas != nil {
		villager.NomorKitas = request.NomorKitas
	}
	if request.NomorPaspor != nil {
		villager.NomorPaspor = request.NomorPaspor
	}

	if err := s.villagerRepo.UpdateVillagerWithTx(tx, villager); err != nil {
		log.Println("Error updating villager:", err)
		return nil, errors.New("failed to update villager")
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}

	return &dtos.MessageResponse{
		Message: "Villager updated successfully",
	}, nil
}

func (s *VillagerService) DeleteVillager(nik *string) (*dtos.MessageResponse, error) {
	tx := s.villagerRepo.BeginTransaction()
	defer tx.Rollback()

	_, err := s.villagerRepo.FindVillagerByNIK(nik)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Villager not found:", err)
			return nil, errors.New("villager not found")
		}
		log.Println("Error finding villager:", err)
		return nil, errors.New("failed to find villager")
	}

	if err := s.villagerRepo.DeleteVillagerWithTx(tx, nik); err != nil {
		log.Println("Error deleting villager:", err)
		return nil, errors.New("failed to delete villager")
	}
	if err := tx.Commit().Error; err != nil {
		log.Println("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}

	return &dtos.MessageResponse{
		Message: "Villager deleted successfully",
	}, nil
}
