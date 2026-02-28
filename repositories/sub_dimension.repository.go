package repositories

import (
	"Apps-I_Desa_Backend/config"
	"Apps-I_Desa_Backend/models"
	"gorm.io/gorm"
)

type SubDimensionRepository struct {
	DB *gorm.DB
}

func NewSubDimensionRepository() *SubDimensionRepository {
	return &SubDimensionRepository{
		DB: config.DB,
	}
}

func (r *SubDimensionRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}

func (r *SubDimensionRepository) CreateSubDimensionPendidikanWithTx(
	tx *gorm.DB,
	pendidikan *models.SubDimensiPendidikan,
) error {
	return tx.Create(pendidikan).Error
}

func (r *SubDimensionRepository) CreateSubDimensionKesehatanWithTx(
	tx *gorm.DB,
	kesehatan *models.SubDimensiKesehatan,
) error {
	return tx.Create(kesehatan).Error
}

func (r *SubDimensionRepository) CreateSubDimensionUtilitasDasarWithTx(
	tx *gorm.DB,
	utilitas *models.SubDimensiUtilitasDasar,
) error {
	return tx.Create(utilitas).Error
}

func (r *SubDimensionRepository) CreateSubDimensionAktivitasWithTx(
	tx *gorm.DB,
	aktivitas *models.SubDimensiAktivitas,
) error {
	return tx.Create(aktivitas).Error
}

func (r *SubDimensionRepository) CreateSubDimensionFasilitasMasyarakatWithTx(
	tx *gorm.DB,
	fasilitas *models.SubDimensiFasilitasMasyarakat,
) error {
	return tx.Create(fasilitas).Error
}

func (r *SubDimensionRepository) CreateSubDimensionProduksiDesaWithTx(
	tx *gorm.DB,
	produksi *models.SubDimensiProduksiDesa,
) error {
	return tx.Create(produksi).Error
}

func (r *SubDimensionRepository) CreateSubDimensionFasilitasPendukungEkonomiWithTx(
	tx *gorm.DB,
	fasilitas *models.SubDimensiFasilitasPendukungEkonomi,
) error {
	return tx.Create(fasilitas).Error
}

func (r *SubDimensionRepository) CreateSubDimensionPengelolaanLingkunganWithTx(
	tx *gorm.DB,
	lingkungan *models.SubDimensiPengelolaanLingkungan,
) error {
	return tx.Create(lingkungan).Error
}

func (r *SubDimensionRepository) CreateSubDimensionPenanggulanganBencanaWithTx(
	tx *gorm.DB,
	bencana *models.SubDimensiPenanggulanganBencana,
) error {
	return tx.Create(bencana).Error
}

func (r *SubDimensionRepository) CreateSubDimensionKondisiAksesJalanWithTx(
	tx *gorm.DB,
	jalan *models.SubDimensiKondisiAksesJalan,
) error {
	return tx.Create(jalan).Error
}

func (r *SubDimensionRepository) CreateSubDimensionKemudahanAksesWithTx(
	tx *gorm.DB,
	akses *models.SubDimensiKemudahanAkses,
) error {
	return tx.Create(akses).Error
}

func (r *SubDimensionRepository) CreateSubDimensionKelembagaanPelayananDesaWithTx(
	tx *gorm.DB,
	kelembagaan *models.SubDimensiKelembagaanPelayananDesa,
) error {
	return tx.Create(kelembagaan).Error
}

func (r *SubDimensionRepository) CreateSubDimensionTataKelolaKeuanganDesaWithTx(
	tx *gorm.DB,
	keuangan *models.SubDimensiTataKelolaKeuanganDesa,
) error {
	return tx.Create(keuangan).Error
}
