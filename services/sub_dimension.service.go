package services

import (
	"errors"

	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/models"
	"Apps-I_Desa_Backend/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type SubDimensionService struct {
	subDimensionRepo *repositories.SubDimensionRepository
}

func NewSubDimensionService(
	subDimensionRepo *repositories.SubDimensionRepository,
) *SubDimensionService {
	return &SubDimensionService{
		subDimensionRepo: subDimensionRepo,
	}
}

func (s *SubDimensionService) CreateSubDimensionPendidikan(
	req *dtos.AddSubDimensionPendidikanRequest,
	ctx *fiber.Ctx,
) (*dtos.MessageResponse, error) {
	tx := s.subDimensionRepo.BeginTransaction()
	defer tx.Rollback()

	villageIDStr := ctx.Locals("village").(string)
	// Parse the village ID from the context
	if villageIDStr == "" {
		log.Error("Village ID not found in context")
		return nil, errors.New("village ID not found")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Error("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID")
	}

	subDimensiPendidikan := &models.SubDimensiPendidikan{
		VillageID:          villageID,
		Year:               *req.Year,
		KetersediaanPaud:   req.KetersediaanPaud,
		KemudahanAksesPaud: req.KemudahanAksesPaud,
		ApmPaud:            req.ApmPaud,
		KemudahanAksesSd:   req.KemudahanAksesSd,
		ApmSd:              req.ApmSd,
		KemudahanAksesSmp:  req.KemudahanAksesSmp,
		ApmSmp:             req.ApmSmp,
		KemudahanAksesSma:  req.KemudahanAksesSma,
		ApmSma:             req.ApmSma,
	}
	if err := s.subDimensionRepo.CreateSubDimensionPendidikanWithTx(tx, subDimensiPendidikan); err != nil {
		log.Error("Error creating sub dimension pendidikan:", err)
		return nil, errors.New("failed to create sub dimension pendidikan")
	}
	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}
	return &dtos.MessageResponse{
		Message: "Sub Dimension Pendidikan created successfully",
	}, nil
}

func (s *SubDimensionService) CreateSubDimensionKesehatan(
	req *dtos.AddSubDimensionKesehatanRequest,
	ctx *fiber.Ctx,
) (*dtos.MessageResponse, error) {
	tx := s.subDimensionRepo.BeginTransaction()
	defer tx.Rollback()

	villageIDStr := ctx.Locals("village").(string)
	// Parse the village ID from the context
	if villageIDStr == "" {
		log.Error("Village ID not found in context")
		return nil, errors.New("village ID not found")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Error("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID")
	}

	subDimensiKesehatan := &models.SubDimensiKesehatan{
		VillageID:                                  villageID,
		Year:                                       *req.Year,
		KemudahanAksesSaranaKesehatan:              req.KemudahanAksesSaranaKesehatan,
		KetersediaanFasilitasKesehatan:             req.KetersediaanFasilitasKesehatan,
		KemudahanAksesFasilitasKesehatan:           req.KemudahanAksesFasilitasKesehatan,
		KetersediaanPosyandu:                       req.KetersediaanPosyandu,
		JumlahAktivitasPosyandu:                    req.JumlahAktivitasPosyandu,
		KemudahanAksesPosyandu:                     req.KemudahanAksesPosyandu,
		KetersediaanLayananDokter:                  req.KetersediaanLayananDokter,
		HariOperasionalLayananDokter:               req.HariOperasionalLayananDokter,
		PenyediaLayananDokter:                      req.PenyediaLayananDokter,
		PenyediaTransportasiLayananDokter:          req.PenyediaTransportasiLayananDokter,
		KetersediaanLayananBidan:                   req.KetersediaanLayananBidan,
		HariOperasionalLayananBidan:                req.HariOperasionalLayananBidan,
		PenyediaLayananBidan:                       req.PenyediaLayananBidan,
		PenyediaTransportasiLayananBidan:           req.PenyediaTransportasiLayananBidan,
		KetersediaanLayananTenagaKesehatan:         req.KetersediaanLayananTenagaKesehatan,
		HariOperasionalLayananTenagaKesehatan:      req.HariOperasionalLayananTenagaKesehatan,
		PenyediaLayananTenagaKesehatan:             req.PenyediaLayananTenagaKesehatan,
		PenyediaTransportasiLayananTenagaKesehatan: req.PenyediaTransportasiLayananTenagaKesehatan,
		PersentasePesertaJaminanKesehatan:          req.PersentasePesertaJaminanKesehatan,
		KegiatanSosialisasiJaminanKesehatan:        req.KegiatanSosialisasiJaminanKesehatan,
	}

	if err := s.subDimensionRepo.CreateSubDimensionKesehatanWithTx(tx, subDimensiKesehatan); err != nil {
		log.Error("Error creating sub dimension kesehatan:", err)
		return nil, errors.New("failed to create sub dimension kesehatan")
	}
	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}
	return &dtos.MessageResponse{
		Message: "Sub Dimension Kesehatan created successfully",
	}, nil
}

func (s *SubDimensionService) CreateSubDimensionUtilitasDasar(
	req *dtos.AddSubDimensionUtilitasDasarRequest,
	ctx *fiber.Ctx,
) (*dtos.MessageResponse, error) {
	tx := s.subDimensionRepo.BeginTransaction()
	defer tx.Rollback()

	villageIDStr := ctx.Locals("village").(string)
	// Parse the village ID from the context
	if villageIDStr == "" {
		log.Error("Village ID not found in context")
		return nil, errors.New("village ID not found")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Error("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID")
	}

	subDimensiUtilitasDasar := &models.SubDimensiUtilitasDasar{
		VillageID:                     villageID,
		Year:                          *req.Year,
		OperasionalAirMinum:           req.OperasionalAirMinum,
		KetersediaanAirMinum:          req.KetersediaanAirMinum,
		KemudahanAksesAirMinum:        req.KemudahanAksesAirMinum,
		KualitasAirMinum:              req.KualitasAirMinum,
		PersentaseRumahTidakLayakHuni: req.PersentaseRumahTidakLayakHuni,
	}
	if err := s.subDimensionRepo.CreateSubDimensionUtilitasDasarWithTx(tx, subDimensiUtilitasDasar); err != nil {
		log.Error("Error creating sub dimension utilitas dasar:", err)
		return nil, errors.New("failed to create sub dimension utilitas dasar")
	}
	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}

	return &dtos.MessageResponse{
		Message: "Sub Dimension Utilitas Dasar created successfully",
	}, nil
}

func (s *SubDimensionService) CreateSubDimensionAktivitas(
	req *dtos.AddSubDimensionAktivitasRequest,
	ctx *fiber.Ctx,
) (*dtos.MessageResponse, error) {
	tx := s.subDimensionRepo.BeginTransaction()
	defer tx.Rollback()

	villageIDStr := ctx.Locals("village").(string)
	// Parse the village ID from the context
	if villageIDStr == "" {
		log.Error("Village ID not found in context")
		return nil, errors.New("village ID not found")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Error("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID")
	}

	subDimensiAktivitas := &models.SubDimensiAktivitas{
		VillageID:                         villageID,
		Year:                              *req.Year,
		KearifanBudayaSosial:              req.KearifanBudayaSosial,
		KearifanBudayaSosialDipertahankan: req.KearifanBudayaSosialDipertahankan,
		KegiatanGotongRoyong:              req.KegiatanGotongRoyong,
		FrekuensiGotongRoyong:             req.FrekuensiGotongRoyong,
		KeterlibatanWargaGotongRoyong:     req.KeterlibatanWargaGotongRoyong,
		FrekuensiKegiatanOlahraga:         req.FrekuensiKegiatanOlahraga,
		PenyelesaianKonflikSecaraDamai:    req.PenyelesaianKonflikSecaraDamai,
		PeranAparatKeamananMediator:       req.PeranAparatKeamananMediator,
		PeranAparatPemerintah:             req.PeranAparatPemerintah,
		PeranTokohMasyarakat:              req.PeranTokohMasyarakat,
		PeranTokohAgama:                   req.PeranTokohAgama,
		SatuanKeamananLingkungan:          req.SatuanKeamananLingkungan,
		AktivitasSatuanKeamananLingkungan: req.AktivitasSatuanKeamananLingkungan,
	}

	if err := s.subDimensionRepo.CreateSubDimensionAktivitasWithTx(tx, subDimensiAktivitas); err != nil {
		log.Error("Error creating sub dimension aktivitas:", err)
		return nil, errors.New("failed to create sub dimension aktivitas")
	}
	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}

	return &dtos.MessageResponse{
		Message: "Sub Dimension Aktivitas created successfully",
	}, nil
}

func (s *SubDimensionService) CreateSubDimensionFasilitasMasyarakat(
	req *dtos.AddSubDimensionFasilitasMasyarakatRequest,
	ctx *fiber.Ctx,
) (*dtos.MessageResponse, error) {
	tx := s.subDimensionRepo.BeginTransaction()
	defer tx.Rollback()

	villageIDStr := ctx.Locals("village").(string)
	// Parse the village ID from the context
	if villageIDStr == "" {
		log.Error("Village ID not found in context")
		return nil, errors.New("village ID not found")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Error("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID")
	}

	subDimensiFasilitasMasyarakat := &models.SubDimensiFasilitasMasyarakat{
		VillageID:                            villageID,
		Year:                                 *req.Year,
		TerdapatTamanBacaanMasyarakat:        req.TerdapatTamanBacaanMasyarakat,
		HariOperasionalTamanBacaanMasyarakat: req.HariOperasionalTamanBacaanMasyarakat,
		KetersediaanFasilitasOlahraga:        req.KetersediaanFasilitasOlahraga,
		KeberadaanRuangPublikTerbuka:         req.KeberadaanRuangPublikTerbuka,
	}
	if err := s.subDimensionRepo.CreateSubDimensionFasilitasMasyarakatWithTx(tx, subDimensiFasilitasMasyarakat); err != nil {
		log.Error("Error creating sub dimension fasilitas masyarakat:", err)
		return nil, errors.New("failed to create sub dimension fasilitas masyarakat")
	}
	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}

	return &dtos.MessageResponse{
		Message: "Sub Dimension Fasilitas Masyarakat created successfully",
	}, nil
}

func (s *SubDimensionService) CreateSubDimensionProduksiDesa(
	req *dtos.AddSubDimensionProduksiDesaRequest,
	ctx *fiber.Ctx,
) (*dtos.MessageResponse, error) {
	tx := s.subDimensionRepo.BeginTransaction()
	defer tx.Rollback()

	villageIDStr := ctx.Locals("village").(string)
	// Parse the village ID from the context
	if villageIDStr == "" {
		log.Error("Village ID not found in context")
		return nil, errors.New("village ID not found")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Error("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID")
	}

	subDimensiPendidikan := &models.SubDimensiProduksiDesa{
		VillageID:                                villageID,
		Year:                                     *req.Year,
		KeragamanAktivitasEkonomi:                req.KeragamanAktivitasEkonomi,
		KeaktifanAktivitasEkonomi:                req.KeaktifanAktivitasEkonomi,
		KetersediaanProdukUnggulanDesa:           req.KetersediaanProdukUnggulanDesa,
		CakupanPasarProdukUnggulan:               req.CakupanPasarProdukUnggulan,
		KetersediaanMerekDagang:                  req.KetersediaanMerekDagang,
		TerdapatKearifanLokalEkonomi:             req.TerdapatKearifanLokalEkonomi,
		TelahDilakukanKerjaSamaDenganDesaLainnya: req.TelahDilakukanKerjaSamaDenganDesaLainnya,
		TelahDilakukanKerjaSamaDenganPihakKetiga: req.TelahDilakukanKerjaSamaDenganPihakKetiga,
	}

	if err := s.subDimensionRepo.CreateSubDimensionProduksiDesaWithTx(tx, subDimensiPendidikan); err != nil {
		log.Error("Error creating sub dimension produksi desa:", err)
		return nil, errors.New("failed to create sub dimension produksi desa")
	}
	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}

	return &dtos.MessageResponse{
		Message: "Sub Dimension Produksi Desa created successfully",
	}, nil
}

func (s *SubDimensionService) CreateSubDimensionFasilitasPendukungEkonomi(
	req *dtos.AddSubDimensionFasilitasPendukungEkonomiRequest,
	ctx *fiber.Ctx,
) (*dtos.MessageResponse, error) {
	tx := s.subDimensionRepo.BeginTransaction()
	defer tx.Rollback()

	villageIDStr := ctx.Locals("village").(string)
	// Parse the village ID from the context
	if villageIDStr == "" {
		log.Error("Village ID not found in context")
		return nil, errors.New("village ID not found")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Error("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID")
	}

	model := &models.SubDimensiFasilitasPendukungEkonomi{
		VillageID:                         villageID,
		Year:                              *req.Year,
		KetersediaanPendidikanNonFormal:   req.KetersediaanPendidikanNonFormal,
		KeterlibatanPendidikanNonFormal:   req.KeterlibatanPendidikanNonFormal,
		KetersediaanPasarRakyat:           req.KetersediaanPasarRakyat,
		KemudahanAksesPasarRakyat:         req.KemudahanAksesPasarRakyat,
		KetersediaanToko:                  req.KetersediaanToko,
		KemudahanAksesToko:                req.KemudahanAksesToko,
		KetersediaanRumahMakan:            req.KetersediaanRumahMakan,
		KemudahanAksesRumahMakan:          req.KemudahanAksesRumahMakan,
		KetersediaanPenginapan:            req.KetersediaanPenginapan,
		KemudahanAksesPenginapan:          req.KemudahanAksesPenginapan,
		KetersediaanLogistik:              req.KetersediaanLogistik,
		KemudahanAksesLogistik:            req.KemudahanAksesLogistik,
		TerdapatBumd:                      req.TerdapatBumd,
		BumdBerbadanHukum:                 req.BumdBerbadanHukum,
		HariOperasionalLembagaEkonomi:     req.HariOperasionalLembagaEkonomi,
		KetersediaanLembagaEkonomiLainnya: req.KetersediaanLembagaEkonomiLainnya,
		KetersediaanKud:                   req.KetersediaanKud,
		KetersediaanUmkm:                  req.KetersediaanUmkm,
		LayananPerbankan:                  req.LayananPerbankan,
		HariOperasionalKeuangan:           req.HariOperasionalKeuangan,
		LayananFasilitasKreditKur:         req.LayananFasilitasKreditKur,
		LayananFasilitasKreditKkpE:        req.LayananFasilitasKreditKkpE,
		LayananFasilitasKreditKuk:         req.LayananFasilitasKreditKuk,
		StatusLayananFasilitasKredit:      req.StatusLayananFasilitasKredit,
	}

	if err := s.subDimensionRepo.CreateSubDimensionFasilitasPendukungEkonomiWithTx(tx, model); err != nil {
		log.Error("Error creating sub dimension fasilitas pendukung ekonomi:", err)
		return nil, errors.New("failed to create sub dimension fasilitas pendukung ekonomi")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}

	return &dtos.MessageResponse{
		Message: "Sub Dimension Fasilitas Pendukung Ekonomi created successfully",
	}, nil
}

func (s *SubDimensionService) CreateSubDimensionPengelolaanLingkungan(
	req *dtos.AddSubDimensionPengelolaanLingkunganRequest,
	ctx *fiber.Ctx,
) (*dtos.MessageResponse, error) {
	tx := s.subDimensionRepo.BeginTransaction()
	defer tx.Rollback()

	villageIDStr := ctx.Locals("village").(string)
	// Parse the village ID from the context
	if villageIDStr == "" {
		log.Error("Village ID not found in context")
		return nil, errors.New("village ID not found")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Error("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID")
	}

	model := &models.SubDimensiPengelolaanLingkungan{
		VillageID:                         villageID,
		Year:                              *req.Year,
		UpayaMenjagaKelestarianLingkungan: req.UpayaMenjagaKelestarianLingkungan,
		RegulasiPelestarianLingkungan:     req.RegulasiPelestarianLingkungan,
		KegiatanPelestarianLingkungan:     req.KegiatanPelestarianLingkungan,
		PemanfaatanEnergiTerbarukan:       req.PemanfaatanEnergiTerbarukan,
		TempatPembuananganSampah:          req.TempatPembuananganSampah,
		PengelolaanSampah:                 req.PengelolaanSampah,
		PemanfaatanSampah:                 req.PemanfaatanSampah,
		KejadianPencemaranLingkungan:      req.KejadianPencemaranLingkungan,
		KetersediaanJamban:                req.KetersediaanJamban,
		KeberfungsianJamban:               req.KeberfungsianJamban,
		KetersediaanSepticTank:            req.KetersediaanSepticTank,
		PembuanganAirLimbahCairRumah:      req.PembuanganAirLimbahCairRumah,
	}

	if err := s.subDimensionRepo.CreateSubDimensionPengelolaanLingkunganWithTx(tx, model); err != nil {
		log.Error("Error creating sub dimension pengelolaan lingkungan:", err)
		return nil, errors.New("failed to create sub dimension pengelolaan lingkungan")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}

	return &dtos.MessageResponse{
		Message: "Sub Dimension Pengelolaan Lingkungan created successfully",
	}, nil
}

func (s *SubDimensionService) CreateSubDimensionPenanggulanganBencana(
	req *dtos.AddSubDimensionPenanggulanganBencanaRequest,
	ctx *fiber.Ctx,
) (*dtos.MessageResponse, error) {
	tx := s.subDimensionRepo.BeginTransaction()
	defer tx.Rollback()

	villageIDStr := ctx.Locals("village").(string)
	// Parse the village ID from the context
	if villageIDStr == "" {
		log.Error("Village ID not found in context")
		return nil, errors.New("village ID not found")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Error("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID")
	}

	model := &models.SubDimensiPenanggulanganBencana{
		VillageID:                           villageID,
		Year:                                *req.Year,
		AspekInformasiKebencanaan:           req.AspekInformasiKebencanaan,
		FasilitasMitigasiBencana:            req.FasilitasMitigasiBencana,
		AksesMenujuFasilitasMitigasiBencana: req.AksesMenujuFasilitasMitigasiBencana,
		AktivitasMitigasi:                   req.AktivitasMitigasi,
		FasilitasTanggapDaruratBencana:      req.FasilitasTanggapDaruratBencana,
	}

	if err := s.subDimensionRepo.CreateSubDimensionPenanggulanganBencanaWithTx(tx, model); err != nil {
		log.Error("Error creating sub dimension penanggulangan bencana:", err)
		return nil, errors.New("failed to create sub dimension penanggulangan bencana")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}

	return &dtos.MessageResponse{
		Message: "Sub Dimension Penanggulangan Bencana created successfully",
	}, nil
}

func (s *SubDimensionService) CreateSubDimensionKondisiAksesJalan(
	req *dtos.AddSubDimensionKondisiAksesJalanRequest,
	ctx *fiber.Ctx,
) (*dtos.MessageResponse, error) {
	tx := s.subDimensionRepo.BeginTransaction()
	defer tx.Rollback()

	villageIDStr := ctx.Locals("village").(string)
	// Parse the village ID from the context
	if villageIDStr == "" {
		log.Error("Village ID not found in context")
		return nil, errors.New("village ID not found")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Error("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID")
	}

	model := &models.SubDimensiKondisiAksesJalan{
		VillageID:            villageID,
		Year:                 *req.Year,
		JenisPermukaanJalan:  req.JenisPermukaanJalan,
		KualitasJalan:        req.KualitasJalan,
		PeneranganJalanUtama: req.PeneranganJalanUtama,
		OperasionalPju:       req.OperasionalPju,
	}

	if err := s.subDimensionRepo.CreateSubDimensionKondisiAksesJalanWithTx(tx, model); err != nil {
		log.Error("Error creating sub dimension kondisi akses jalan:", err)
		return nil, errors.New("failed to create sub dimension kondisi akses jalan")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}

	return &dtos.MessageResponse{
		Message: "Sub Dimension Kondisi Akses Jalan created successfully",
	}, nil
}

func (s *SubDimensionService) CreateSubDimensionKemudahanAkses(
	req *dtos.AddSubDimensionKemudahanAksesRequest,
	ctx *fiber.Ctx,
) (*dtos.MessageResponse, error) {
	tx := s.subDimensionRepo.BeginTransaction()
	defer tx.Rollback()

	villageIDStr := ctx.Locals("village").(string)
	// Parse the village ID from the context
	if villageIDStr == "" {
		log.Error("Village ID not found in context")
		return nil, errors.New("village ID not found")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Error("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID")
	}

	model := &models.SubDimensiKemudahanAkses{
		VillageID:                    villageID,
		Year:                         *req.Year,
		AngkutanPerdesaan:            req.AngkutanPerdesaan,
		OperasionalAngkutanPerdesaan: req.OperasionalAngkutanPerdesaan,
		PelayananListrik:             req.PelayananListrik,
		DurasiLayananListrik:         req.DurasiLayananListrik,
		AksesTelepon:                 req.AksesTelepon,
		AksesInternet:                req.AksesInternet,
	}

	if err := s.subDimensionRepo.CreateSubDimensionKemudahanAksesWithTx(tx, model); err != nil {
		log.Error("Error creating sub dimension kemudahan akses:", err)
		return nil, errors.New("failed to create sub dimension kemudahan akses")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}

	return &dtos.MessageResponse{
		Message: "Sub Dimension Kemudahan Akses created successfully",
	}, nil
}

func (s *SubDimensionService) CreateSubDimensionKelembagaanPelayananDesa(
	req *dtos.AddSubDimensionKelembagaanPelayananDesaRequest,
	ctx *fiber.Ctx,
) (*dtos.MessageResponse, error) {
	tx := s.subDimensionRepo.BeginTransaction()
	defer tx.Rollback()

	villageIDStr := ctx.Locals("village").(string)
	// Parse the village ID from the context
	if villageIDStr == "" {
		log.Error("Village ID not found in context")
		return nil, errors.New("village ID not found")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Error("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID")
	}

	model := &models.SubDimensiKelembagaanPelayananDesa{
		VillageID:                              villageID,
		Year:                                   *req.Year,
		LayananDiberikan:                       req.LayananDiberikan,
		PublikasiInformasiPelayanan:            req.PublikasiInformasiPelayanan,
		PelayananAdministrasi:                  req.PelayananAdministrasi,
		PelayananPengaduan:                     req.PelayananPengaduan,
		PelayananLainnya:                       req.PelayananLainnya,
		MusyawarahDesa:                         req.MusyawarahDesa,
		MusyawarahDesaDidatangiUnsurMasyarakat: req.MusyawarahDesaDidatangiUnsurMasyarakat,
	}

	if err := s.subDimensionRepo.CreateSubDimensionKelembagaanPelayananDesaWithTx(tx, model); err != nil {
		log.Error("Error creating sub dimension kelembagaan pelayanan desa:", err)
		return nil, errors.New("failed to create sub dimension kelembagaan pelayanan desa")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}

	return &dtos.MessageResponse{
		Message: "Sub Dimension Kelembagaan Pelayanan Desa created successfully",
	}, nil
}

func (s *SubDimensionService) CreateSubDimensionTataKelolaKeuanganDesa(
	req *dtos.AddSubDimensionTataKelolaKeuanganDesaRequest,
	ctx *fiber.Ctx,
) (*dtos.MessageResponse, error) {
	tx := s.subDimensionRepo.BeginTransaction()
	defer tx.Rollback()

	villageIDStr := ctx.Locals("village").(string)
	// Parse the village ID from the context
	if villageIDStr == "" {
		log.Error("Village ID not found in context")
		return nil, errors.New("village ID not found")
	}
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Error("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID")
	}

	model := &models.SubDimensiTataKelolaKeuanganDesa{
		VillageID:             villageID,
		Year:                  *req.Year,
		PendapatanAsliDesa:    req.PendapatanAsliDesa,
		PeningkatanPades:      req.PeningkatanPades,
		PenyertaanModalDdBumd: req.PenyertaanModalDdBumd,
		AsetTanahDesa:         req.AsetTanahDesa,
		AsetKantorDesa:        req.AsetKantorDesa,
		AsetPasarDesa:         req.AsetPasarDesa,
		AsetLainnya:           req.AsetLainnya,
		ProduktivitasAsetDesa: req.ProduktivitasAsetDesa,
		InventarisasiAsetDesa: req.InventarisasiAsetDesa,
	}

	if err := s.subDimensionRepo.CreateSubDimensionTataKelolaKeuanganDesaWithTx(tx, model); err != nil {
		log.Error("Error creating sub dimension tata kelola keuangan desa:", err)
		return nil, errors.New("failed to create sub dimension tata kelola keuangan desa")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}

	return &dtos.MessageResponse{
		Message: "Sub Dimension Tata Kelola Keuangan Desa created successfully",
	}, nil
}
