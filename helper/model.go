package helper

import (
	"context"
	"minang-kos-service/constant"
	"minang-kos-service/model/domain"
	"minang-kos-service/model/web/dto"
	"minang-kos-service/model/web/response"
	"time"

	"github.com/golang-jwt/jwt"
)

func BuildBaseDomain(ctx context.Context) domain.BaseDomain {
	userInfo := GetUserInfo(ctx)

	return domain.BaseDomain{
		CreatedAt:     time.Now(),
		CreatedBy:     userInfo.UserId,
		CreatedByName: userInfo.Name,
		UpdatedAt:     time.Now(),
		UpdatedBy:     userInfo.UserId,
		UpdatedByName: userInfo.Name,
		IsDeleted:     false,
	}
}

func SetUpdatedBy(ctx context.Context, baseDomain *domain.BaseDomain) {
	userInfo := GetUserInfo(ctx)

	baseDomain.UpdatedAt = time.Now()
	baseDomain.UpdatedBy = userInfo.UserId
	baseDomain.UpdatedByName = userInfo.Name
}

func GetUserInfo(ctx context.Context) dto.UserInfo {
	user := ctx.Value(constant.KeyContext(constant.USER_INFO)).(jwt.MapClaims)
	return dto.UserInfo{
		UserId: int64(user["userId"].(float64)),
		Email:  user["email"].(string),
		Name:   user["name"].(string),
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

func ToCityResponse(city domain.City) response.CityResponse {
	return response.CityResponse{
		Id:                 city.Id,
		Name:               city.Name,
		Province:           ToProvinceResponse(city.Province),
		BaseDomainResponse: ToBaseDomainResponse(city.BaseDomain),
	}
}

func ToCityResponses(cities []domain.City) []response.CityResponse {
	if cities == nil {
		return make([]response.CityResponse, 0)
	}

	var cityResponses []response.CityResponse
	for _, city := range cities {
		cityResponses = append(cityResponses, ToCityResponse(city))
	}
	return cityResponses
}

func ToRoleResponse(role domain.Role) response.RoleResponse {
	return response.RoleResponse{
		Id:                 role.Id,
		Name:               role.Name,
		BaseDomainResponse: ToBaseDomainResponse(role.BaseDomain),
	}
}

func ToRoleResponses(roles []domain.Role) []response.RoleResponse {
	if roles == nil {
		return make([]response.RoleResponse, 0)
	}

	var roleResponses []response.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, ToRoleResponse(role))
	}
	return roleResponses
}

func ToLoginResponse(token string, expiredAt int64) response.AuthLoginResponse {
	return response.AuthLoginResponse{
		Token:     token,
		ExpiredAt: expiredAt,
	}
}
