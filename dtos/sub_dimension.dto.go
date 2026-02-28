package dtos

type AddSubDimensionPendidikanRequest struct {
	KetersediaanPaud   string `json:"ketersediaan_paud"    validate:"required"`
	KemudahanAksesPaud string `json:"kemudahan_akses_paud" validate:"required"`
	ApmPaud            string `json:"apm_paud"             validate:"required"`
	KemudahanAksesSd   string `json:"kemudahan_akses_sd"   validate:"required"`
	ApmSd              string `json:"apm_sd"               validate:"required"`
	KemudahanAksesSmp  string `json:"kemudahan_akses_smp"  validate:"required"`
	ApmSmp             string `json:"apm_smp"              validate:"required"`
	KemudahanAksesSma  string `json:"kemudahan_akses_sma"  validate:"required"`
	ApmSma             string `json:"apm_sma"              validate:"required"`
	Year               *int   `json:"year"                 validate:"omitempty,gte=2000,lte=2100"` // Assuming year is between 2000 and 2100
}

type AddSubDimensionKesehatanRequest struct {
	KemudahanAksesSaranaKesehatan              string `json:"kemudahan_akses_sarana_kesehatan"               validate:"required"`
	KetersediaanFasilitasKesehatan             string `json:"ketersediaan_fasilitas_kesehatan"               validate:"required"`
	KemudahanAksesFasilitasKesehatan           string `json:"kemudahan_akses_fasilitas_kesehatan"            validate:"required"`
	KetersediaanPosyandu                       string `json:"ketersediaan_posyandu"                          validate:"required"`
	JumlahAktivitasPosyandu                    string `json:"jumlah_aktivitas_posyandu"                      validate:"required"`
	KemudahanAksesPosyandu                     string `json:"kemudahan_akses_posyandu"                       validate:"required"`
	KetersediaanLayananDokter                  string `json:"ketersediaan_layanan_dokter"                    validate:"required"`
	HariOperasionalLayananDokter               string `json:"hari_operasional_layanan_dokter"                validate:"required"`
	PenyediaLayananDokter                      string `json:"penyedia_layanan_dokter"                        validate:"required"`
	PenyediaTransportasiLayananDokter          string `json:"penyedia_transportasi_layanan_dokter"           validate:"required"`
	KetersediaanLayananBidan                   string `json:"ketersediaan_layanan_bidan"                     validate:"required"`
	HariOperasionalLayananBidan                string `json:"hari_operasional_layanan_bidan"                 validate:"required"`
	PenyediaLayananBidan                       string `json:"penyedia_layanan_bidan"                         validate:"required"`
	PenyediaTransportasiLayananBidan           string `json:"penyedia_transportasi_layanan_bidan"            validate:"required"`
	KetersediaanLayananTenagaKesehatan         string `json:"ketersediaan_layanan_tenaga_kesehatan"          validate:"required"`
	HariOperasionalLayananTenagaKesehatan      string `json:"hari_operasional_layanan_tenaga_kesehatan"      validate:"required"`
	PenyediaLayananTenagaKesehatan             string `json:"penyedia_layanan_tenaga_kesehatan"              validate:"required"`
	PenyediaTransportasiLayananTenagaKesehatan string `json:"penyedia_transportasi_layanan_tenaga_kesehatan" validate:"required"`
	PersentasePesertaJaminanKesehatan          string `json:"persentase_peserta_jaminan_kesehatan"           validate:"required"`
	KegiatanSosialisasiJaminanKesehatan        string `json:"kegiatan_sosialisasi_jaminan_kesehatan"         validate:"required"`
	Year                                       *int   `json:"year"                                           validate:"omitempty,gte=2000,lte=2100"` // Assuming year is
}

type AddSubDimensionUtilitasDasarRequest struct {
	OperasionalAirMinum           string `json:"operasional_air_minum"             validate:"required"`
	KetersediaanAirMinum          string `json:"ketersediaan_air_minum"            validate:"required"`
	KemudahanAksesAirMinum        string `json:"kemudahan_akses_air_minum"         validate:"required"`
	KualitasAirMinum              string `json:"kualitas_air_minum"                validate:"required"`
	PersentaseRumahTidakLayakHuni string `json:"persentase_rumah_tidak_layak_huni" validate:"required"`
	Year                          *int   `json:"year"                              validate:"omitempty,gte=2000,lte=2100"` // Assuming year is between 2000 and 2100
}

type AddSubDimensionAktivitasRequest struct {
	KearifanBudayaSosial              string `json:"kearifan_budaya_sosial"               validate:"required"`
	KearifanBudayaSosialDipertahankan string `json:"kearifan_budaya_sosial_dipertahankan" validate:"required"`
	KegiatanGotongRoyong              string `json:"kegiatan_gotong_royong"               validate:"required"`
	FrekuensiGotongRoyong             string `json:"frekuensi_gotong_royong"              validate:"required"`
	KeterlibatanWargaGotongRoyong     string `json:"keterlibatan_warga_gotong_royong"     validate:"required"`
	FrekuensiKegiatanOlahraga         string `json:"frekuensi_kegiatan_olahraga"          validate:"required"`
	PenyelesaianKonflikSecaraDamai    string `json:"penyelesaian_konflik_secara_damai"    validate:"required"`
	PeranAparatKeamananMediator       string `json:"peran_aparat_keamanan_mediator"       validate:"required"`
	PeranAparatPemerintah             string `json:"peran_aparat_pemerintah"              validate:"required"`
	PeranTokohMasyarakat              string `json:"peran_tokoh_masyarakat"               validate:"required"`
	PeranTokohAgama                   string `json:"peran_tokoh_agama"                    validate:"required"`
	SatuanKeamananLingkungan          string `json:"satuan_keamanan_lingkungan"           validate:"required"`
	AktivitasSatuanKeamananLingkungan string `json:"aktivitas_satuan_keamanan_lingkungan" validate:"required"`
	Year                              *int   `json:"year"                                 validate:"omitempty,gte=2000,lte=2100"` // Assuming year is between 2000 and 2100
}

type AddSubDimensionFasilitasMasyarakatRequest struct {
	TerdapatTamanBacaanMasyarakat        string `json:"terdapat_taman_bacaan_masyarakat"         validate:"required"`
	HariOperasionalTamanBacaanMasyarakat string `json:"hari_operasional_taman_bacaan_masyarakat" validate:"required"`
	KetersediaanFasilitasOlahraga        string `json:"ketersediaan_fasilitas_olahraga"          validate:"required"`
	KeberadaanRuangPublikTerbuka         string `json:"keberadaan_ruang_publik_terbuka"          validate:"required"`
	Year                                 *int   `json:"year"                                     validate:"omitempty,gte=2000,lte=2100"`
}

type AddSubDimensionProduksiDesaRequest struct {
	KeragamanAktivitasEkonomi                string `json:"keragaman_aktivitas_ekonomi"                    validate:"required"`
	KeaktifanAktivitasEkonomi                string `json:"keaktifan_aktivitas_ekonomi"                    validate:"required"`
	KetersediaanProdukUnggulanDesa           string `json:"ketersediaan_produk_unggulan_desa"              validate:"required"`
	CakupanPasarProdukUnggulan               string `json:"cakupan_pasar_produk_unggulan"                  validate:"required"`
	KetersediaanMerekDagang                  string `json:"ketersediaan_merek_dagang"                      validate:"required"`
	TerdapatKearifanLokalEkonomi             string `json:"terdapat_kearifan_lokal_ekonomi"                validate:"required"`
	TelahDilakukanKerjaSamaDenganDesaLainnya string `json:"telah_dilakukan_kerja_sama_dengan_desa_lainnya" validate:"required"`
	TelahDilakukanKerjaSamaDenganPihakKetiga string `json:"telah_dilakukan_kerja_sama_dengan_pihak_ketiga" validate:"required"`
	Year                                     *int   `json:"year"                                           validate:"omitempty,gte=2000,lte=2100"`
}

type AddSubDimensionFasilitasPendukungEkonomiRequest struct {
	KetersediaanPendidikanNonFormal   string `json:"ketersediaan_pendidikan_non_formal"   validate:"required"`
	KeterlibatanPendidikanNonFormal   string `json:"keterlibatan_pendidikan_non_formal"   validate:"required"`
	KetersediaanPasarRakyat           string `json:"ketersediaan_pasar_rakyat"            validate:"required"`
	KemudahanAksesPasarRakyat         string `json:"kemudahan_akses_pasar_rakyat"         validate:"required"`
	KetersediaanToko                  string `json:"ketersediaan_toko"                    validate:"required"`
	KemudahanAksesToko                string `json:"kemudahan_akses_toko"                 validate:"required"`
	KetersediaanRumahMakan            string `json:"ketersediaan_rumah_makan"             validate:"required"`
	KemudahanAksesRumahMakan          string `json:"kemudahan_akses_rumah_makan"          validate:"required"`
	KetersediaanPenginapan            string `json:"ketersediaan_penginapan"              validate:"required"`
	KemudahanAksesPenginapan          string `json:"kemudahan_akses_penginapan"           validate:"required"`
	KetersediaanLogistik              string `json:"ketersediaan_logistik"                validate:"required"`
	KemudahanAksesLogistik            string `json:"kemudahan_akses_logistik"             validate:"required"`
	TerdapatBumd                      string `json:"terdapat_bumd"                        validate:"required"`
	BumdBerbadanHukum                 string `json:"bumd_berbadan_hukum"                  validate:"required"`
	HariOperasionalLembagaEkonomi     string `json:"hari_operasional_lembaga_ekonomi"     validate:"required"`
	KetersediaanLembagaEkonomiLainnya string `json:"ketersediaan_lembaga_ekonomi_lainnya" validate:"required"`
	KetersediaanKud                   string `json:"ketersediaan_kud"                     validate:"required"`
	KetersediaanUmkm                  string `json:"ketersediaan_umkm"                    validate:"required"`
	LayananPerbankan                  string `json:"layanan_perbankan"                    validate:"required"`
	HariOperasionalKeuangan           string `json:"hari_operasional_keuangan"            validate:"required"`
	LayananFasilitasKreditKur         string `json:"layanan_fasilitas_kredit_kur"         validate:"required"`
	LayananFasilitasKreditKkpE        string `json:"layanan_fasilitas_kredit_kkp_e"       validate:"required"`
	LayananFasilitasKreditKuk         string `json:"layanan_fasilitas_kredit_kuk"         validate:"required"`
	StatusLayananFasilitasKredit      string `json:"status_layanan_fasilitas_kredit"      validate:"required"`
	Year                              *int   `json:"year"                                 validate:"omitempty,gte=2000,lte=2100"`
}

type AddSubDimensionPengelolaanLingkunganRequest struct {
	UpayaMenjagaKelestarianLingkungan string `json:"upaya_menjaga_kelestarian_lingkungan" validate:"required"`
	RegulasiPelestarianLingkungan     string `json:"regulasi_pelestarian_lingkungan"      validate:"required"`
	KegiatanPelestarianLingkungan     string `json:"kegiatan_pelestarian_lingkungan"      validate:"required"`
	PemanfaatanEnergiTerbarukan       string `json:"pemanfaatan_energi_terbarukan"        validate:"required"`
	TempatPembuananganSampah          string `json:"tempat_pembuanangan_sampah"           validate:"required"`
	PengelolaanSampah                 string `json:"pengelolaan_sampah"                   validate:"required"`
	PemanfaatanSampah                 string `json:"pemanfaatan_sampah"                   validate:"required"`
	KejadianPencemaranLingkungan      string `json:"kejadian_pencemaran_lingkungan"       validate:"required"`
	KetersediaanJamban                string `json:"ketersediaan_jamban"                  validate:"required"`
	KeberfungsianJamban               string `json:"keberfungsian_jamban"                 validate:"required"`
	KetersediaanSepticTank            string `json:"ketersediaan_septic_tank"             validate:"required"`
	PembuanganAirLimbahCairRumah      string `json:"pembuangan_air_limbah_cair_rumah"     validate:"required"`
	Year                              *int   `json:"year"                                 validate:"omitempty,gte=2000,lte=2100"`
}

type AddSubDimensionPenanggulanganBencanaRequest struct {
	AspekInformasiKebencanaan           string `json:"aspek_informasi_kebencanaan"             validate:"required"`
	FasilitasMitigasiBencana            string `json:"fasilitas_mitigasi_bencana"              validate:"required"`
	AksesMenujuFasilitasMitigasiBencana string `json:"akses_menuju_fasilitas_mitigasi_bencana" validate:"required"`
	AktivitasMitigasi                   string `json:"aktivitas_mitigasi"                      validate:"required"`
	FasilitasTanggapDaruratBencana      string `json:"fasilitas_tanggap_darurat_bencana"       validate:"required"`
	Year                                *int   `json:"year"                                    validate:"omitempty,gte=2000,lte=2100"`
}

type AddSubDimensionKondisiAksesJalanRequest struct {
	JenisPermukaanJalan  string `json:"jenis_permukaan_jalan"  validate:"required"`
	KualitasJalan        string `json:"kualitas_jalan"         validate:"required"`
	PeneranganJalanUtama string `json:"penerangan_jalan_utama" validate:"required"`
	OperasionalPju       string `json:"operasional_pju"        validate:"required"`
	Year                 *int   `json:"year"                   validate:"omitempty,gte=2000,lte=2100"`
}

type AddSubDimensionKemudahanAksesRequest struct {
	AngkutanPerdesaan            string `json:"angkutan_perdesaan"             validate:"required"`
	OperasionalAngkutanPerdesaan string `json:"operasional_angkutan_perdesaan" validate:"required"`
	PelayananListrik             string `json:"pelayanan_listrik"              validate:"required"`
	DurasiLayananListrik         string `json:"durasi_layanan_listrik"         validate:"required"`
	AksesTelepon                 string `json:"akses_telepon"                  validate:"required"`
	AksesInternet                string `json:"akses_internet"                 validate:"required"`
	Year                         *int   `json:"year"                           validate:"omitempty,gte=2000,lte=2100"`
}

type AddSubDimensionKelembagaanPelayananDesaRequest struct {
	LayananDiberikan                       string `json:"layanan_diberikan"                          validate:"required"`
	PublikasiInformasiPelayanan            string `json:"publikasi_informasi_pelayanan"              validate:"required"`
	PelayananAdministrasi                  string `json:"pelayanan_administrasi"                     validate:"required"`
	PelayananPengaduan                     string `json:"pelayanan_pengaduan"                        validate:"required"`
	PelayananLainnya                       string `json:"pelayanan_lainnya"                          validate:"required"`
	MusyawarahDesa                         string `json:"musyawarah_desa"                            validate:"required"`
	MusyawarahDesaDidatangiUnsurMasyarakat string `json:"musyawarah_desa_didatangi_unsur_masyarakat" validate:"required"`
	Year                                   *int   `json:"year"                                       validate:"omitempty,gte=2000,lte=2100"`
}

type AddSubDimensionTataKelolaKeuanganDesaRequest struct {
	PendapatanAsliDesa    string `json:"pendapatan_asli_desa"     validate:"required"`
	PeningkatanPades      string `json:"peningkatan_pades"        validate:"required"`
	PenyertaanModalDdBumd string `json:"penyertaan_modal_dd_bumd" validate:"required"`
	AsetTanahDesa         string `json:"aset_tanah_desa"          validate:"required"`
	AsetKantorDesa        string `json:"aset_kantor_desa"         validate:"required"`
	AsetPasarDesa         string `json:"aset_pasar_desa"          validate:"required"`
	AsetLainnya           string `json:"aset_lainnya"             validate:"required"`
	ProduktivitasAsetDesa string `json:"produktivitas_aset_desa"  validate:"required"`
	InventarisasiAsetDesa string `json:"inventarisasi_aset_desa"  validate:"required"`
	Year                  *int   `json:"year"                     validate:"omitempty,gte=2000,lte=2100"`
}
