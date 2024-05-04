package request

type FacilityUpdateRequest struct {
	Id             int64  `validate:"required"`
	Icon           string `validate:"required,max=255,min=1" json:"icon"`
	Name           string `validate:"required,max=255,min=1" json:"name"`
	FacilityTypeId int64  `validate:"required" json:"facilityTypeId"`
}
