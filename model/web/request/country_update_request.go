package request

type CountryUpdateRequest struct {
	Id   int64  `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}
