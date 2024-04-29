package request

type ProvinceCreateRequest struct {
	Name      string `validate:"required,max=200,min=1" json:"name"`
	CountryId int64  `validate:"required" json:"countryId"`
}
