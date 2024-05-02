package request

type FacilityCreateRequest struct {
	Name           string `validate:"required,max=255,min=1" json:"name"`
	FacilityTypeId int64  `validate:"required" json:"facilityTypeId"`
}
