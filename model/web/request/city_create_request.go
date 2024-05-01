package request

type CityCreateRequest struct {
	Name       string `validate:"required,max=255,min=1" json:"name"`
	ProvinceId int64  `validate:"required" json:"provinceId"`
}
