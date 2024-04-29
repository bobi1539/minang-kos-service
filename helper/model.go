package helper

import (
	"minang-kos-service/model/domain"
	"minang-kos-service/model/web/response"
)

func ToResponsePagination(page int, size int, totalItem int, data any) response.PaginationResponse {
	return response.PaginationResponse{
		Page:      page,
		Size:      size,
		TotalPage: (totalItem / size) + 1,
		TotalItem: totalItem,
		Data:      data,
	}
}

func ToCountryResponse(country domain.Country) response.CountryResponse {
	return response.CountryResponse{
		Id:            country.Id,
		Name:          country.Name,
		CreatedAt:     country.CreatedAt,
		CreatedBy:     country.CreatedBy,
		CreatedByName: country.CreatedByName,
		UpdatedAt:     country.UpdatedAt,
		UpdatedBy:     country.UpdatedBy,
		UpdatedByName: country.UpdatedByName,
		IsDeleted:     country.IsDeleted,
	}
}

func ToCountryResponses(countries []domain.Country) []response.CountryResponse {
	if countries == nil {
		return make([]response.CountryResponse, 0)
	}

	var countryResponses []response.CountryResponse
	for _, country := range countries {
		countryResponses = append(countryResponses, ToCountryResponse(country))
	}
	return countryResponses
}
