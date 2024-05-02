package request

type FacilityTypeUpdateRequest struct {
	Id   int64  `validate:"required"`
	Name string `validate:"required,max=255,min=1" json:"name"`
}
