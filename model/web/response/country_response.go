package response

import "time"

type CountryResponse struct {
	Id            int64     `json:"id"`
	Name          string    `json:"Name"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     int64     `json:"createdBy"`
	CreatedByName string    `json:"createdByName"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UpdatedBy     int64     `json:"updatedBy"`
	UpdatedByName string    `json:"updatedByName"`
	IsDeleted     bool      `json:"isDeleted"`
}
