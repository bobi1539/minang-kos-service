package controller

import (
	"errors"
	"minang-kos-service/constant"
	"minang-kos-service/exception"
	"minang-kos-service/helper"
	"minang-kos-service/model/web/request"
	"minang-kos-service/model/web/search"
	"minang-kos-service/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const PROVINCE_ID = "provinceId"
const PROVINCE_NAME = "name"

type ProvinceControllerImpl struct {
	ProvinceService service.ProvinceService
}

func NewProvinceController(provinceService service.ProvinceService) ProvinceController {
	return &ProvinceControllerImpl{
		ProvinceService: provinceService,
	}
}

func (controller *ProvinceControllerImpl) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	provinceRequest := request.ProvinceCreateRequest{}
	helper.ReadFromRequestBody(httpRequest, &provinceRequest)

	provinceResponse := controller.ProvinceService.Create(httpRequest.Context(), provinceRequest)
	helper.WriteSuccessResponse(writer, provinceResponse)
}

func (controller *ProvinceControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	provinceRequest := request.ProvinceUpdateRequest{}
	helper.ReadFromRequestBody(httpRequest, &provinceRequest)

	provinceRequest.Id = helper.GetIdFromPath(params, PROVINCE_ID)
	provinceResponse := controller.ProvinceService.Update(httpRequest.Context(), provinceRequest)
	helper.WriteSuccessResponse(writer, provinceResponse)
}

func (controller *ProvinceControllerImpl) Delete(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	provinceId := helper.GetIdFromPath(params, PROVINCE_ID)
	controller.ProvinceService.Delete(httpRequest.Context(), provinceId)
	helper.WriteSuccessResponse(writer, nil)
}

func (controller *ProvinceControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	provinceId := helper.GetIdFromPath(params, PROVINCE_ID)
	provinceResponse := controller.ProvinceService.FindById(httpRequest.Context(), provinceId)
	helper.WriteSuccessResponse(writer, provinceResponse)
}

func (controller *ProvinceControllerImpl) FindAllWithPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := search.ProvinceSearch{
		Name:      helper.GetQueryParam(httpRequest, PROVINCE_NAME),
		CountryId: helper.StringToInt64(helper.GetQueryParam(httpRequest, COUNTRY_ID)),
		PageSize:  search.BuildPageSizeFromRequest(httpRequest),
	}

	provinceResponses := controller.ProvinceService.FindAllWithPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, provinceResponses)
}

func (controller *ProvinceControllerImpl) FindAllWithoutPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := search.ProvinceSearch{
		Name:      helper.GetQueryParam(httpRequest, PROVINCE_NAME),
		CountryId: getCountryId(helper.GetQueryParam(httpRequest, COUNTRY_ID)),
		PageSize:  search.BuildPageSize(0, 0),
	}

	provinceResponses := controller.ProvinceService.FindAllWithoutPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, provinceResponses)
}

func getCountryId(countryId string) int64 {
	if len(countryId) == 0 {
		exception.PanicErrorBadRequest(errors.New(constant.COUNTRY_ID_REQUIRED))
	}
	return helper.StringToInt64(countryId)
}
