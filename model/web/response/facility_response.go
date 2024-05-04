package response

type FacilityResponse struct {
	Id           int64                `json:"id"`
	Name         string               `json:"name"`
	Icon         string               `json:"icon"`
	FacilityType FacilityTypeResponse `json:"facilityType"`
	BaseDomainResponse
}
