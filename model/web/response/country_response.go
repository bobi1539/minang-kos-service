package response

type CountryResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"Name"`
	BaseDomainResponse
}
