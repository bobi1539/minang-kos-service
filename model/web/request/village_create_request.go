package request

type VillageCreateRequest struct {
	Name       string `validate:"required,max=255,min=1" json:"name"`
	DistrictId int64  `validate:"required" json:"districtId"`
}
