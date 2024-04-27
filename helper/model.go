package helper

import (
	"minang-kos-service/model/domain"
	"minang-kos-service/model/web/response"
)

func ToCountryResponse(country domain.Country) response.CountryResponse {
	return response.CountryResponse{
		Id:            country.Id,
		Name:          country.Name,
		CreatedAt:     country.UpdatedAt,
		CreatedBy:     country.CreatedBy,
		CreatedByName: country.UpdatedByName,
		UpdatedAt:     country.UpdatedAt,
		UpdatedBy:     country.UpdatedBy,
		UpdatedByName: country.UpdatedByName,
		IsDeleted:     country.IsDeleted,
	}
}
