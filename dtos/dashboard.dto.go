package dtos

type GetDashboardResponse struct {
	TotalKeluarga       int32   `json:"totalKeluarga"`
	TotalPenduduk       int32   `json:"totalPenduduk"`
	RerataKeluarga      float32 `json:"rerataKeluarga"`
	TotalLakiLaki       int32   `json:"lakiLaki"`
	TotalPerempuan      int32   `json:"perempuan"`
	TotalKepalaKeluarga int32   `json:"kepalaKeluarga"`
	RerataUmur          float32 `json:"rerataUmur"`
	TotalRT             int32   `json:"rt"`
	TotalRW             int32   `json:"rw"`
	TotalKelurahan      int32   `json:"kelurahan"`
	TotalKecamatan      int32   `json:"kecamatan"`
}
