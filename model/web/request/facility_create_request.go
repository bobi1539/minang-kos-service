package request

type FacilityCreateRequest struct {
	Name           string `validate:"required,max=255,min=1" json:"name"`
	Icon           string `validate:"required,max=255,min=1" json:"icon"`
	FacilityTypeId int64  `validate:"required" json:"facilityTypeId"`
}
