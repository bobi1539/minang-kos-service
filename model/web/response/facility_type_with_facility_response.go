package response

type FacilityTypeWithFacilityResponse struct {
	Id         int64              `json:"id"`
	Name       string             `json:"name"`
	Facilities []FacilityResponse `json:"facilities"`
}
