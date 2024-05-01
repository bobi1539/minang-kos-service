package request

type DistrictCreateRequest struct {
	Name   string `validate:"required,max=255,min=1" json:"name"`
	CityId int64  `validate:"required" json:"cityId"`
}
