package request

type CityUpdateRequest struct {
	Id         int64  `validate:"required"`
	Name       string `validate:"required,max=255,min=1" json:"name"`
	ProvinceId int64  `validate:"required" json:"provinceId"`
}
