package request

type FacilityTypeCreateRequest struct {
	Name string `validate:"required,max=255,min=1" json:"name"`
}
