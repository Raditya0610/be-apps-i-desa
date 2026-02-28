package dtos

// FamilyCardsResponse represents the response structure for family cards.
type GetAllFamilyCardsResponse struct {
	FamilyCards []GetFamilyCardResponse `json:"family_cards"`
}

// FamilyCardResponse represents the structure of a single family card response.
type GetFamilyCardResponse struct {
	NIK          string  `json:"nik"`
	Name         *string `json:"name,omitempty"`
	TotalMembers int     `json:"total_members"`
}

type GetAllFamilyMember struct {
	NIK           string            `json:"nik"`
	Address       string            `json:"address"`
	FamilyMembers []GetFamilyMember `json:"family_members"`
}

type AddFamilyCardRequest struct {
	NIK           string `json:"nik"            validate:"required,len=16,numeric"`
	Address       string `json:"address"        validate:"required"`
	RT            string `json:"rt"             validate:"required"`
	RW            string `json:"rw"             validate:"required"`
	Kelurahan     string `json:"kelurahan"      validate:"required"`
	Kecamatan     string `json:"kecamatan"      validate:"required"`
	KabupatenKota string `json:"kabupaten_kota" validate:"required"`
	KodePos       string `json:"kode_pos"       validate:"required"`
	Provinsi      string `json:"provinsi"       validate:"required"`
}
