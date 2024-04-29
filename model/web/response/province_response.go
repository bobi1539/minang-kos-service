package response

type ProvinceResponse struct {
	Id      int64           `json:"id"`
	Name    string          `json:"Name"`
	Country CountryResponse `json:"country"`
	BaseDomainResponse
}
