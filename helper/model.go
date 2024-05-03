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

func ToDistrictResponse(district domain.District) response.DistrictResponse {
	return response.DistrictResponse{
		Id:                 district.Id,
		Name:               district.Name,
		City:               ToCityResponse(district.City),
		BaseDomainResponse: ToBaseDomainResponse(district.BaseDomain),
	}
}

func ToDistrictResponses(districts []domain.District) []response.DistrictResponse {
	if districts == nil {
		return make([]response.DistrictResponse, 0)
	}

	var districtResponses []response.DistrictResponse
	for _, district := range districts {
		districtResponses = append(districtResponses, ToDistrictResponse(district))
	}
	return districtResponses
}

func ToVillageResponse(village domain.Village) response.VillageResponse {
	return response.VillageResponse{
		Id:                 village.Id,
		Name:               village.Name,
		District:           ToDistrictResponse(village.District),
		BaseDomainResponse: ToBaseDomainResponse(village.BaseDomain),
	}
}

func ToVillageResponses(villages []domain.Village) []response.VillageResponse {
	if villages == nil {
		return make([]response.VillageResponse, 0)
	}

	var villageResponses []response.VillageResponse
	for _, village := range villages {
		villageResponses = append(villageResponses, ToVillageResponse(village))
	}
	return villageResponses
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

func ToKosTypeResponse(kosType domain.KosType) response.KosTypeResponse {
	return response.KosTypeResponse{
		Id:                 kosType.Id,
		Name:               kosType.Name,
		BaseDomainResponse: ToBaseDomainResponse(kosType.BaseDomain),
	}
}

func ToKosTypeResponses(kosTypes []domain.KosType) []response.KosTypeResponse {
	if kosTypes == nil {
		return make([]response.KosTypeResponse, 0)
	}

	var kosTypeResponses []response.KosTypeResponse
	for _, kosType := range kosTypes {
		kosTypeResponses = append(kosTypeResponses, ToKosTypeResponse(kosType))
	}
	return kosTypeResponses
}

func ToFacilityTypeResponse(facilityType domain.FacilityType) response.FacilityTypeResponse {
	return response.FacilityTypeResponse{
		Id:                 facilityType.Id,
		Name:               facilityType.Name,
		BaseDomainResponse: ToBaseDomainResponse(facilityType.BaseDomain),
	}
}

func ToFacilityTypeResponses(facilityTypes []domain.FacilityType) []response.FacilityTypeResponse {
	if facilityTypes == nil {
		return make([]response.FacilityTypeResponse, 0)
	}

	var facilityTypeResponses []response.FacilityTypeResponse
	for _, facilityType := range facilityTypes {
		facilityTypeResponses = append(facilityTypeResponses, ToFacilityTypeResponse(facilityType))
	}
	return facilityTypeResponses
}

func ToFacilityResponse(facility domain.Facility) response.FacilityResponse {
	return response.FacilityResponse{
		Id:                 facility.Id,
		Name:               facility.Name,
		FacilityType:       ToFacilityTypeResponse(facility.FacilityType),
		BaseDomainResponse: ToBaseDomainResponse(facility.BaseDomain),
	}
}

func ToFacilityResponses(facilities []domain.Facility) []response.FacilityResponse {
	if facilities == nil {
		return make([]response.FacilityResponse, 0)
	}

	var facilityResponses []response.FacilityResponse
	for _, facility := range facilities {
		facilityResponses = append(facilityResponses, ToFacilityResponse(facility))
	}
	return facilityResponses
}

func ToUserResponse(user domain.User) response.UserResponse {
	return response.UserResponse{
		Id:                 user.Id,
		Email:              user.Email,
		Name:               user.Name,
		PhoneNumber:        user.PhoneNumber,
		Role:               ToRoleResponse(user.Role),
		BaseDomainResponse: ToBaseDomainResponse(user.BaseDomain),
	}
}

func ToKosBedroomResponse(
	kosBedroom domain.KosBedroom,
	facilityTypes []domain.FacilityType,
	facilities []domain.Facility,
) response.KosBedroomResponse {
	return response.KosBedroomResponse{
		Id:                   kosBedroom.Id,
		Title:                kosBedroom.Title,
		RoomLength:           kosBedroom.RoomLength,
		RoomWidth:            kosBedroom.RoomWidth,
		UnitLength:           kosBedroom.UnitLength,
		IsIncludeElectricity: kosBedroom.IsIncludeElectricity,
		Price:                kosBedroom.Price,
		Images:               nil,
		KosType:              ToKosTypeResponse(kosBedroom.KosType),
		Village:              ToVillageResponse(kosBedroom.Village),
		User:                 ToUserResponse(kosBedroom.User),
		FacilityTypes:        ToFacilityTypeWithFacilityResponses(facilityTypes, facilities),
		BaseDomainResponse:   ToBaseDomainResponse(kosBedroom.BaseDomain),
	}
}

func ToFacilityTypeWithFacilityResponse(facilityType domain.FacilityType, facilities []domain.Facility) response.FacilityTypeWithFacilityResponse {
	var facilityResponses []response.FacilityResponse

	for _, facility := range facilities {
		if facility.FacilityType.Id == facilityType.Id {
			facilityResponses = append(facilityResponses, ToFacilityResponse(facility))
		}
	}

	return response.FacilityTypeWithFacilityResponse{
		Id:         facilityType.Id,
		Name:       facilityType.Name,
		Facilities: facilityResponses,
	}
}

func ToFacilityTypeWithFacilityResponses(facilityTypes []domain.FacilityType, facilities []domain.Facility) []response.FacilityTypeWithFacilityResponse {
	if facilityTypes == nil {
		return make([]response.FacilityTypeWithFacilityResponse, 0)
	}

	var responses []response.FacilityTypeWithFacilityResponse
	for _, facilityType := range facilityTypes {
		response := ToFacilityTypeWithFacilityResponse(facilityType, facilities)
		responses = append(responses, response)
	}
	return responses
}
