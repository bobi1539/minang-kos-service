package helper

import (
	"minang-kos-service/model/domain"
	"minang-kos-service/model/web/response"
	"time"
)

func BuildBaseDomain() domain.BaseDomain {
	return domain.BaseDomain{
		CreatedAt:     time.Now(),
		CreatedBy:     1,
		CreatedByName: "test",
		UpdatedAt:     time.Now(),
		UpdatedBy:     1,
		UpdatedByName: "test",
		IsDeleted:     false,
	}
}

func ToResponsePagination(page int, size int, totalItem int, data any) response.PaginationResponse {
	return response.PaginationResponse{
		Page:      page,
		Size:      size,
		TotalPage: (totalItem / size) + 1,
		TotalItem: totalItem,
		Data:      data,
	}
}

func ToBaseDomainResponse(baseDomain domain.BaseDomain) response.BaseDomainResponse {
	return response.BaseDomainResponse{
		CreatedAt:     baseDomain.CreatedAt,
		CreatedBy:     baseDomain.CreatedBy,
		CreatedByName: baseDomain.CreatedByName,
		UpdatedAt:     baseDomain.UpdatedAt,
		UpdatedBy:     baseDomain.UpdatedBy,
		UpdatedByName: baseDomain.UpdatedByName,
		IsDeleted:     baseDomain.IsDeleted,
	}
}

func ToCountryResponse(country domain.Country) response.CountryResponse {
	return response.CountryResponse{
		Id:                 country.Id,
		Name:               country.Name,
		BaseDomainResponse: ToBaseDomainResponse(country.BaseDomain),
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

func ToProvinceResponse(province domain.Province) response.ProvinceResponse {
	return response.ProvinceResponse{
		Id:                 province.Id,
		Name:               province.Name,
		Country:            ToCountryResponse(province.Country),
		BaseDomainResponse: ToBaseDomainResponse(province.BaseDomain),
	}
}

func ToProvinceResponses(provinces []domain.Province) []response.ProvinceResponse {
	if provinces == nil {
		return make([]response.ProvinceResponse, 0)
	}

	var provinceResponses []response.ProvinceResponse
	for _, province := range provinces {
		provinceResponses = append(provinceResponses, ToProvinceResponse(province))
	}
	return provinceResponses
}
