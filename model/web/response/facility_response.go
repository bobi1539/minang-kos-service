package response

type FacilityResponse struct {
	Id           int64                `json:"id"`
	Name         string               `json:"name"`
	FacilityType FacilityTypeResponse `json:"facilityType"`
	BaseDomainResponse
}
