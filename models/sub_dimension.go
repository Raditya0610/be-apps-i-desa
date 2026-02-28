package models

import "github.com/google/uuid"

type SubDimensiPendidikan struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID          uuid.UUID `gorm:"type:uuid;not null"`
	Year               int       `gorm:"size:4"`
	KetersediaanPaud   string    `gorm:"size:50;not null"`
	KemudahanAksesPaud string    `gorm:"size:50;not null"`
	ApmPaud            string    `gorm:"size:50;not null"`
	KemudahanAksesSd   string    `gorm:"size:50;not null"`
	ApmSd              string    `gorm:"size:50;not null"`
	KemudahanAksesSmp  string    `gorm:"size:50;not null"`
	ApmSmp             string    `gorm:"size:50;not null"`
	KemudahanAksesSma  string    `gorm:"size:50;not null"`
	ApmSma             string    `gorm:"size:50;not null"`
	Village            Village   `gorm:"foreignKey:VillageID"`
}

type SubDimensiKesehatan struct {
	ID                                         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                                  uuid.UUID `gorm:"type:uuid;not null"`
	Year                                       int       `gorm:"size:4"`
	KemudahanAksesSaranaKesehatan              string    `gorm:"size:50;not null"`
	KetersediaanFasilitasKesehatan             string    `gorm:"size:50;not null"`
	KemudahanAksesFasilitasKesehatan           string    `gorm:"size:50;not null"`
	KetersediaanPosyandu                       string    `gorm:"size:50;not null"`
	JumlahAktivitasPosyandu                    string    `gorm:"size:50;not null"`
	KemudahanAksesPosyandu                     string    `gorm:"size:50;not null"`
	KetersediaanLayananDokter                  string    `gorm:"size:50;not null"`
	HariOperasionalLayananDokter               string    `gorm:"size:50;not null"`
	PenyediaLayananDokter                      string    `gorm:"size:50;not null"`
	PenyediaTransportasiLayananDokter          string    `gorm:"size:50;not null"`
	KetersediaanLayananBidan                   string    `gorm:"size:50;not null"`
	HariOperasionalLayananBidan                string    `gorm:"size:50;not null"`
	PenyediaLayananBidan                       string    `gorm:"size:50;not null"`
	PenyediaTransportasiLayananBidan           string    `gorm:"size:50;not null"`
	KetersediaanLayananTenagaKesehatan         string    `gorm:"size:50;not null"`
	HariOperasionalLayananTenagaKesehatan      string    `gorm:"size:50;not null"`
	PenyediaLayananTenagaKesehatan             string    `gorm:"size:50;not null"`
	PenyediaTransportasiLayananTenagaKesehatan string    `gorm:"size:50;not null"`
	PersentasePesertaJaminanKesehatan          string    `gorm:"size:50;not null"`
	KegiatanSosialisasiJaminanKesehatan        string    `gorm:"size:50;not null"`
	Village                                    Village   `gorm:"foreignKey:VillageID"`
}

type SubDimensiUtilitasDasar struct {
	ID                            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                     uuid.UUID `gorm:"type:uuid;not null"`
	Year                          int       `gorm:"size:4"`
	OperasionalAirMinum           string    `gorm:"size:50;not null"`
	KetersediaanAirMinum          string    `gorm:"size:50;not null"`
	KemudahanAksesAirMinum        string    `gorm:"size:50;not null"`
	KualitasAirMinum              string    `gorm:"size:50;not null"`
	PersentaseRumahTidakLayakHuni string    `gorm:"size:50;not null"`
	Village                       Village   `gorm:"foreignKey:VillageID"`
}

type SubDimensiAktivitas struct {
	ID                                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                         uuid.UUID `gorm:"type:uuid;not null"`
	Year                              int       `gorm:"size:4"`
	KearifanBudayaSosial              string    `gorm:"size:50;not null"`
	KearifanBudayaSosialDipertahankan string    `gorm:"size:50;not null"`
	KegiatanGotongRoyong              string    `gorm:"size:50;not null"`
	FrekuensiGotongRoyong             string    `gorm:"size:50;not null"`
	KeterlibatanWargaGotongRoyong     string    `gorm:"size:50;not null"`
	FrekuensiKegiatanOlahraga         string    `gorm:"size:50;not null"`
	PenyelesaianKonflikSecaraDamai    string    `gorm:"size:50;not null"`
	PeranAparatKeamananMediator       string    `gorm:"size:50;not null"`
	PeranAparatPemerintah             string    `gorm:"size:50;not null"`
	PeranTokohMasyarakat              string    `gorm:"size:50;not null"`
	PeranTokohAgama                   string    `gorm:"size:50;not null"`
	SatuanKeamananLingkungan          string    `gorm:"size:50;not null"`
	AktivitasSatuanKeamananLingkungan string    `gorm:"size:50;not null"`
	Village                           Village   `gorm:"foreignKey:VillageID"`
}

type SubDimensiFasilitasMasyarakat struct {
	ID                                   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                            uuid.UUID `gorm:"type:uuid;not null"`
	Year                                 int       `gorm:"size:4"`
	TerdapatTamanBacaanMasyarakat        string    `gorm:"size:50;not null"`
	HariOperasionalTamanBacaanMasyarakat string    `gorm:"size:50;not null"`
	KetersediaanFasilitasOlahraga        string    `gorm:"size:50;not null"`
	KeberadaanRuangPublikTerbuka         string    `gorm:"size:50;not null"`
	Village                              Village   `gorm:"foreignKey:VillageID"`
}

type SubDimensiProduksiDesa struct {
	ID                                       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                                uuid.UUID `gorm:"type:uuid;not null"`
	Year                                     int       `gorm:"size:4"`
	KeragamanAktivitasEkonomi                string    `gorm:"size:50;not null"`
	KeaktifanAktivitasEkonomi                string    `gorm:"size:50;not null"`
	KetersediaanProdukUnggulanDesa           string    `gorm:"size:50;not null"`
	CakupanPasarProdukUnggulan               string    `gorm:"size:50;not null"`
	KetersediaanMerekDagang                  string    `gorm:"size:50;not null"`
	TerdapatKearifanLokalEkonomi             string    `gorm:"size:50;not null"`
	TelahDilakukanKerjaSamaDenganDesaLainnya string    `gorm:"size:50;not null"`
	TelahDilakukanKerjaSamaDenganPihakKetiga string    `gorm:"size:50;not null"`
	Village                                  Village   `gorm:"foreignKey:VillageID"`
}

type SubDimensiFasilitasPendukungEkonomi struct {
	ID                                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                         uuid.UUID `gorm:"type:uuid;not null"`
	Year                              int       `gorm:"size:4"`
	KetersediaanPendidikanNonFormal   string    `gorm:"size:50;not null"`
	KeterlibatanPendidikanNonFormal   string    `gorm:"size:50;not null"`
	KetersediaanPasarRakyat           string    `gorm:"size:50;not null"`
	KemudahanAksesPasarRakyat         string    `gorm:"size:50;not null"`
	KetersediaanToko                  string    `gorm:"size:50;not null"`
	KemudahanAksesToko                string    `gorm:"size:50;not null"`
	KetersediaanRumahMakan            string    `gorm:"size:50;not null"`
	KemudahanAksesRumahMakan          string    `gorm:"size:50;not null"`
	KetersediaanPenginapan            string    `gorm:"size:50;not null"`
	KemudahanAksesPenginapan          string    `gorm:"size:50;not null"`
	KetersediaanLogistik              string    `gorm:"size:50;not null"`
	KemudahanAksesLogistik            string    `gorm:"size:50;not null"`
	TerdapatBumd                      string    `gorm:"size:50;not null"`
	BumdBerbadanHukum                 string    `gorm:"size:50;not null"`
	HariOperasionalLembagaEkonomi     string    `gorm:"size:50;not null"`
	KetersediaanLembagaEkonomiLainnya string    `gorm:"size:50;not null"`
	KetersediaanKud                   string    `gorm:"size:50;not null"`
	KetersediaanUmkm                  string    `gorm:"size:50;not null"`
	LayananPerbankan                  string    `gorm:"size:50;not null"`
	HariOperasionalKeuangan           string    `gorm:"size:50;not null"`
	LayananFasilitasKreditKur         string    `gorm:"size:50;not null"`
	LayananFasilitasKreditKkpE        string    `gorm:"size:50;not null"`
	LayananFasilitasKreditKuk         string    `gorm:"size:50;not null"`
	StatusLayananFasilitasKredit      string    `gorm:"size:50;not null"`
	Village                           Village   `gorm:"foreignKey:VillageID"`
}

type SubDimensiPengelolaanLingkungan struct {
	ID                                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                         uuid.UUID `gorm:"type:uuid;not null"`
	Year                              int       `gorm:"size:4"`
	UpayaMenjagaKelestarianLingkungan string    `gorm:"size:50;not null"`
	RegulasiPelestarianLingkungan     string    `gorm:"size:50;not null"`
	KegiatanPelestarianLingkungan     string    `gorm:"size:50;not null"`
	PemanfaatanEnergiTerbarukan       string    `gorm:"size:50;not null"`
	TempatPembuananganSampah          string    `gorm:"size:50;not null"`
	PengelolaanSampah                 string    `gorm:"size:50;not null"`
	PemanfaatanSampah                 string    `gorm:"size:50;not null"`
	KejadianPencemaranLingkungan      string    `gorm:"size:50;not null"`
	KetersediaanJamban                string    `gorm:"size:50;not null"`
	KeberfungsianJamban               string    `gorm:"size:50;not null"`
	KetersediaanSepticTank            string    `gorm:"size:50;not null"`
	PembuanganAirLimbahCairRumah      string    `gorm:"size:50;not null"`
	Village                           Village   `gorm:"foreignKey:VillageID"`
}

type SubDimensiPenanggulanganBencana struct {
	ID                                  uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                           uuid.UUID `gorm:"type:uuid;not null"`
	Year                                int       `gorm:"size:4"`
	AspekInformasiKebencanaan           string    `gorm:"size:50;not null"`
	FasilitasMitigasiBencana            string    `gorm:"size:50;not null"`
	AksesMenujuFasilitasMitigasiBencana string    `gorm:"size:50;not null"`
	AktivitasMitigasi                   string    `gorm:"size:50;not null"`
	FasilitasTanggapDaruratBencana      string    `gorm:"size:50;not null"`
	Village                             Village   `gorm:"foreignKey:VillageID"`
}

type SubDimensiKondisiAksesJalan struct {
	ID                   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID            uuid.UUID `gorm:"type:uuid;not null"`
	Year                 int       `gorm:"size:4"`
	JenisPermukaanJalan  string    `gorm:"size:50;not null"`
	KualitasJalan        string    `gorm:"size:50;not null"`
	PeneranganJalanUtama string    `gorm:"size:50;not null"`
	OperasionalPju       string    `gorm:"size:50;not null"`
	Village              Village   `gorm:"foreignKey:VillageID"`
}

type SubDimensiKemudahanAkses struct {
	ID                           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                    uuid.UUID `gorm:"type:uuid;not null"`
	Year                         int       `gorm:"size:4"`
	AngkutanPerdesaan            string    `gorm:"size:50;not null"`
	OperasionalAngkutanPerdesaan string    `gorm:"size:50;not null"`
	PelayananListrik             string    `gorm:"size:50;not null"`
	DurasiLayananListrik         string    `gorm:"size:50;not null"`
	AksesTelepon                 string    `gorm:"size:50;not null"`
	AksesInternet                string    `gorm:"size:50;not null"`
	Village                      Village   `gorm:"foreignKey:VillageID"`
}

type SubDimensiKelembagaanPelayananDesa struct {
	ID                                     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                              uuid.UUID `gorm:"type:uuid;not null"`
	Year                                   int       `gorm:"size:4"`
	LayananDiberikan                       string    `gorm:"size:50;not null"`
	PublikasiInformasiPelayanan            string    `gorm:"size:50;not null"`
	PelayananAdministrasi                  string    `gorm:"size:50;not null"`
	PelayananPengaduan                     string    `gorm:"size:50;not null"`
	PelayananLainnya                       string    `gorm:"size:50;not null"`
	MusyawarahDesa                         string    `gorm:"size:50;not null"`
	MusyawarahDesaDidatangiUnsurMasyarakat string    `gorm:"size:50;not null"`
	Village                                Village   `gorm:"foreignKey:VillageID"`
}

type SubDimensiTataKelolaKeuanganDesa struct {
	ID                    uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID             uuid.UUID `gorm:"type:uuid;not null"`
	Year                  int       `gorm:"size:4"`
	PendapatanAsliDesa    string    `gorm:"size:50;not null"`
	PeningkatanPades      string    `gorm:"size:50;not null"`
	PenyertaanModalDdBumd string    `gorm:"size:50;not null"`
	AsetTanahDesa         string    `gorm:"size:50;not null"`
	AsetKantorDesa        string    `gorm:"size:50;not null"`
	AsetPasarDesa         string    `gorm:"size:50;not null"`
	AsetLainnya           string    `gorm:"size:50;not null"`
	ProduktivitasAsetDesa string    `gorm:"size:50;not null"`
	InventarisasiAsetDesa string    `gorm:"size:50;not null"`
	Village               Village   `gorm:"foreignKey:VillageID"`
}
