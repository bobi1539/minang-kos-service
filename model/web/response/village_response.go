package response

type VillageResponse struct {
	Id       int64            `json:"id"`
	Name     string           `json:"name"`
	District DistrictResponse `json:"district"`
	BaseDomainResponse
}
