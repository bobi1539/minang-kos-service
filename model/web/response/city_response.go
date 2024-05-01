package response

type CityResponse struct {
	Id       int64            `json:"id"`
	Name     string           `json:"name"`
	Province ProvinceResponse `json:"province"`
	BaseDomainResponse
}
