package request

type ProvinceUpdateRequest struct {
	Id        int64  `validate:"required"`
	Name      string `validate:"required,max=255,min=1" json:"name"`
	CountryId int64  `validate:"required" json:"countryId"`
}
