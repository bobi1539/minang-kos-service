package response

type DistrictResponse struct {
	Id   int64        `json:"id"`
	Name string       `json:"name"`
	City CityResponse `json:"city"`
	BaseDomainResponse
}
