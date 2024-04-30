package response

type ProvinceResponse struct {
	Id      int64           `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country"`
	BaseDomainResponse
}
