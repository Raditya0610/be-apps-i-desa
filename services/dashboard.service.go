package services

import (
	"errors"
	"sync"

	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type DashboardService struct {
	villagerRepo   *repositories.VillagerRepository
	familyCardRepo *repositories.FamilyCardRepository
}

func NewDashboardService(
	villagerRepo *repositories.VillagerRepository,
	familyCardRepo *repositories.FamilyCardRepository,
) *DashboardService {
	return &DashboardService{
		villagerRepo:   villagerRepo,
		familyCardRepo: familyCardRepo,
	}
}

func (s *DashboardService) GetDashboardData(ctx *fiber.Ctx) (*dtos.GetDashboardResponse, error) {
	villageIDStr := ctx.Locals("village").(string)
	if villageIDStr == "" {
		log.Error("village ID is empty")
		return nil, errors.New("village ID is empty")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Error("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID format")
	}

	// Create channels to receive results
	type result struct {
		value interface{}
		err   error
	}

	// Create channels for each operation
	familyCardsCh := make(chan result, 1)
	rtCh := make(chan result, 1)
	rwCh := make(chan result, 1)
	kelurahanCh := make(chan result, 1)
	kecamatanCh := make(chan result, 1)
	villagersCh := make(chan result, 1)
	maleVillagersCh := make(chan result, 1)
	averageAgeCh := make(chan result, 1)
	kepalaKeluargaCh := make(chan result, 1)

	var wg sync.WaitGroup

	// Launch goroutines for family card operations
	wg.Add(5)
	go func() {
		defer wg.Done()
		count, err := s.familyCardRepo.CountAllFamilyCardByVillageID(&villageID)
		familyCardsCh <- result{count, err}
	}()

	go func() {
		defer wg.Done()
		count, err := s.familyCardRepo.CountDistinctRT(&villageID)
		rtCh <- result{count, err}
	}()

	go func() {
		defer wg.Done()
		count, err := s.familyCardRepo.CountDistinctRW(&villageID)
		rwCh <- result{count, err}
	}()

	go func() {
		defer wg.Done()
		count, err := s.familyCardRepo.CountDistinctKelurahan(&villageID)
		kelurahanCh <- result{count, err}
	}()

	go func() {
		defer wg.Done()
		count, err := s.familyCardRepo.CountDistinctKecamatan(&villageID)
		kecamatanCh <- result{count, err}
	}()

	// Launch goroutines for villager operations
	wg.Add(4)
	go func() {
		defer wg.Done()
		count, err := s.villagerRepo.CountAllVillagerByVillageID(&villageID)
		villagersCh <- result{count, err}
	}()

	go func() {
		defer wg.Done()
		count, err := s.villagerRepo.CountAllLakiLakiVillager(&villageID)
		maleVillagersCh <- result{count, err}
	}()

	go func() {
		defer wg.Done()
		avg, err := s.villagerRepo.GetAverageAge(&villageID)
		averageAgeCh <- result{avg, err}
	}()

	go func() {
		defer wg.Done()
		count, err := s.villagerRepo.CountAllKepalaKeluarga(&villageID)
		kepalaKeluargaCh <- result{count, err}
	}()

	// Wait for all goroutines to complete
	wg.Wait()

	// Collect results and check for errors
	familyCardsRes := <-familyCardsCh
	if familyCardsRes.err != nil {
		log.Error("Error counting family cards:", familyCardsRes.err)
		return nil, errors.New("error counting family cards")
	}
	countFamilyCards := familyCardsRes.value.(int64)

	rtRes := <-rtCh
	if rtRes.err != nil {
		log.Error("Error counting distinct RT:", rtRes.err)
		return nil, errors.New("error counting distinct RT")
	}
	countDistinctRT := rtRes.value.(int64)

	rwRes := <-rwCh
	if rwRes.err != nil {
		log.Error("Error counting distinct RW:", rwRes.err)
		return nil, errors.New("error counting distinct RW")
	}
	countDistinctRW := rwRes.value.(int64)

	kelurahanRes := <-kelurahanCh
	if kelurahanRes.err != nil {
		log.Error("Error counting distinct Kelurahan:", kelurahanRes.err)
		return nil, errors.New("error counting distinct Kelurahan")
	}
	countDistinctKelurahan := kelurahanRes.value.(int64)

	kecamatanRes := <-kecamatanCh
	if kecamatanRes.err != nil {
		log.Error("Error counting distinct Kecamatan:", kecamatanRes.err)
		return nil, errors.New("error counting distinct Kecamatan")
	}
	countDistinctKecamatan := kecamatanRes.value.(int64)

	villagersRes := <-villagersCh
	if villagersRes.err != nil {
		log.Error("Error counting villagers:", villagersRes.err)
		return nil, errors.New("error counting villagers")
	}
	countVillagers := villagersRes.value.(int64)

	maleVillagersRes := <-maleVillagersCh
	if maleVillagersRes.err != nil {
		log.Error("Error counting male villagers:", maleVillagersRes.err)
		return nil, errors.New("error counting male villagers")
	}
	countMaleVillagers := maleVillagersRes.value.(int64)

	averageAgeRes := <-averageAgeCh
	if averageAgeRes.err != nil {
		log.Error("Error getting average age:", averageAgeRes.err)
		return nil, errors.New("error getting average age")
	}
	countAverageAge := averageAgeRes.value.(float32)

	kepalaKeluargaRes := <-kepalaKeluargaCh
	if kepalaKeluargaRes.err != nil {
		log.Error("Error counting kepala keluarga:", kepalaKeluargaRes.err)
		return nil, errors.New("error counting kepala keluarga")
	}
	countKepalaKeluarga := kepalaKeluargaRes.value.(int64)

	return &dtos.GetDashboardResponse{
		TotalKeluarga:       int32(countFamilyCards),
		TotalPenduduk:       int32(countVillagers),
		RerataKeluarga:      float32(countFamilyCards) / float32(countVillagers),
		TotalLakiLaki:       int32(countMaleVillagers),
		TotalPerempuan:      int32(countVillagers) - int32(countMaleVillagers),
		TotalKepalaKeluarga: int32(countKepalaKeluarga),
		RerataUmur:          countAverageAge,
		TotalRT:             int32(countDistinctRT),
		TotalRW:             int32(countDistinctRW),
		TotalKelurahan:      int32(countDistinctKelurahan),
		TotalKecamatan:      int32(countDistinctKecamatan),
	}, nil
}
