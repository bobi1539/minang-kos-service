package controller

import (
	"errors"
	"minang-kos-service/constant"
	"minang-kos-service/exception"
	"minang-kos-service/helper"
	"minang-kos-service/model/web/request"
	"minang-kos-service/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const CITY_ID = "cityId"
const CITY_NAME = "name"

type CityControllerImpl struct {
	CityService service.CityService
}

func NewCityController(cityService service.CityService) CityController {
	return &CityControllerImpl{
		CityService: cityService,
	}
}

func (controller *CityControllerImpl) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	cityRequest := request.CityCreateRequest{}
	helper.ReadFromRequestBody(httpRequest, &cityRequest)

	cityResponse := controller.CityService.Create(httpRequest.Context(), cityRequest)
	helper.WriteSuccessResponse(writer, cityResponse)
}

func (controller *CityControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	cityRequest := request.CityUpdateRequest{}
	helper.ReadFromRequestBody(httpRequest, &cityRequest)

	cityRequest.Id = helper.GetIdFromPath(params, CITY_ID)
	cityResponse := controller.CityService.Update(httpRequest.Context(), cityRequest)
	helper.WriteSuccessResponse(writer, cityResponse)
}

func (controller *CityControllerImpl) Delete(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	cityId := helper.GetIdFromPath(params, CITY_ID)
	controller.CityService.Delete(httpRequest.Context(), cityId)
	helper.WriteSuccessResponse(writer, nil)
}

func (controller *CityControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	cityId := helper.GetIdFromPath(params, CITY_ID)
	cityResponse := controller.CityService.FindById(httpRequest.Context(), cityId)
	helper.WriteSuccessResponse(writer, cityResponse)
}

func (controller *CityControllerImpl) FindAllWithPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := make(map[string]any)
	searchBy["name"] = helper.GetQueryParam(httpRequest, CITY_NAME)
	searchBy["provinceId"] = helper.StringToInt64(helper.GetQueryParam(httpRequest, PROVINCE_ID))
	searchBy["page"] = helper.GetPageOrSize(httpRequest, constant.PAGE)
	searchBy["size"] = helper.GetPageOrSize(httpRequest, constant.SIZE)

	cityResponses := controller.CityService.FindAllWithPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, cityResponses)
}

func (controller *CityControllerImpl) FindAllWithoutPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := make(map[string]any)
	searchBy["name"] = helper.GetQueryParam(httpRequest, CITY_NAME)
	searchBy["provinceId"] = getProvinceId(helper.GetQueryParam(httpRequest, PROVINCE_ID))

	cityResponses := controller.CityService.FindAllWithoutPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, cityResponses)
}

func getProvinceId(provinceId string) int64 {
	if len(provinceId) == 0 {
		exception.PanicErrorBadRequest(errors.New(constant.PROVINCE_ID_REQUIRED))
	}
	return helper.StringToInt64(provinceId)
}
