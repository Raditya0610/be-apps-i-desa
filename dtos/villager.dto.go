package dtos

type AddVillagerRequest struct {
	NIK              string  `json:"nik"               validate:"required,len=16,numeric"`
	NamaLengkap      string  `json:"nama_lengkap"      validate:"required"`
	JenisKelamin     string  `json:"jenis_kelamin"     validate:"required"`
	TempatLahir      string  `json:"tempat_lahir"      validate:"required"`
	TanggalLahir     string  `json:"tanggal_lahir"     validate:"required"`
	Agama            string  `json:"agama"             validate:"required"`
	Pendidikan       string  `json:"pendidikan"        validate:"required"`
	Pekerjaan        string  `json:"pekerjaan"         validate:"required"`
	StatusPerkawinan string  `json:"status_perkawinan" validate:"required"`
	StatusHubungan   string  `json:"status_hubungan"   validate:"required"`
	Kewarganegaraan  string  `json:"kewarganegaraan"   validate:"required"`
	NomorPaspor      *string `json:"nomor_paspor"      validate:"omitempty"`
	NomorKitas       *string `json:"nomor_kitas"       validate:"omitempty"`
	NamaAyah         string  `json:"nama_ayah"         validate:"required"`
	NamaIbu          string  `json:"nama_ibu"          validate:"required"`
	FamilyCardID     string  `json:"family_card_id"    validate:"required"`
}

type GetVillagerResponse struct {
	NIK              string  `json:"nik"`
	NamaLengkap      string  `json:"nama_lengkap"`
	JenisKelamin     string  `json:"jenis_kelamin"`
	TempatLahir      string  `json:"tempat_lahir"`
	TanggalLahir     string  `json:"tanggal_lahir"`
	Agama            string  `json:"agama"`
	Pendidikan       string  `json:"pendidikan"`
	Pekerjaan        string  `json:"pekerjaan"`
	StatusPerkawinan string  `json:"status_perkawinan"`
	StatusHubungan   string  `json:"status_hubungan"`
	Kewarganegaraan  string  `json:"kewarganegaraan"`
	NomorPaspor      *string `json:"nomor_paspor,omitempty"`
	NomorKitas       *string `json:"nomor_kitas,omitempty"`
	NamaAyah         string  `json:"nama_ayah"`
	NamaIbu          string  `json:"nama_ibu"`
	FamilyCardID     string  `json:"family_card_id"`
	VillageID        string  `json:"village_id"`
}

type GetFamilyMember struct {
	NIK            string `json:"nik"`
	Name           string `json:"name"`
	StatusHubungan string `json:"status_hubungan"`
	Age            int    `json:"age"`
	JenisKelamin   string `json:"jenis_kelamin"`
	Pendidikan     string `json:"pendidikan"`
	Pekerjaan      string `json:"pekerjaan"`
}

type UpdateVillagerRequest struct {
	NIK              *string `json:"nik,omitempty"               validate:"omitempty,len=16,numeric"`
	NamaLengkap      *string `json:"nama_lengkap,omitempty"      validate:"omitempty"`
	JenisKelamin     *string `json:"jenis_kelamin,omitempty"     validate:"omitempty"`
	TempatLahir      *string `json:"tempat_lahir,omitempty"      validate:"omitempty"`
	TanggalLahir     *string `json:"tanggal_lahir,omitempty"     validate:"omitempty"`
	Agama            *string `json:"agama,omitempty"             validate:"omitempty"`
	Pendidikan       *string `json:"pendidikan,omitempty"        validate:"omitempty"`
	Pekerjaan        *string `json:"pekerjaan,omitempty"         validate:"omitempty"`
	StatusPerkawinan *string `json:"status_perkawinan,omitempty" validate:"omitempty"`
	StatusHubungan   *string `json:"status_hubungan,omitempty"   validate:"omitempty"`
	Kewarganegaraan  *string `json:"kewarganegaraan,omitempty"   validate:"omitempty"`
	NomorPaspor      *string `json:"nomor_paspor,omitempty"      validate:"omitempty"`
	NomorKitas       *string `json:"nomor_kitas,omitempty"       validate:"omitempty"`
	NamaAyah         *string `json:"nama_ayah,omitempty"         validate:"omitempty"`
	NamaIbu          *string `json:"nama_ibu,omitempty"          validate:"omitempty"`
}

type GetAllVillagersResponse struct {
	NIK  string `json:"nik"`
	Name string `json:"name" gorm:"column:nama_lengkap"`
}
