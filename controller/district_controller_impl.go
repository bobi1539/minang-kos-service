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

const DISTRICT_ID = "districtId"
const DISTRICT_NAME = "name"

type DistrictControllerImpl struct {
	DistrictService service.DistrictService
}

func NewDistrictController(districtService service.DistrictService) DistrictController {
	return &DistrictControllerImpl{
		DistrictService: districtService,
	}
}

func (controller *DistrictControllerImpl) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	districtRequest := request.DistrictCreateRequest{}
	helper.ReadFromRequestBody(httpRequest, &districtRequest)

	districtResponse := controller.DistrictService.Create(httpRequest.Context(), districtRequest)
	helper.WriteSuccessResponse(writer, districtResponse)
}

func (controller *DistrictControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	districtRequest := request.DistrictUpdateRequest{}
	helper.ReadFromRequestBody(httpRequest, &districtRequest)

	districtRequest.Id = helper.GetIdFromPath(params, DISTRICT_ID)
	districtResponse := controller.DistrictService.Update(httpRequest.Context(), districtRequest)
	helper.WriteSuccessResponse(writer, districtResponse)
}

func (controller *DistrictControllerImpl) Delete(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	districtId := helper.GetIdFromPath(params, DISTRICT_ID)
	controller.DistrictService.Delete(httpRequest.Context(), districtId)
	helper.WriteSuccessResponse(writer, nil)
}

func (controller *DistrictControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	districtId := helper.GetIdFromPath(params, DISTRICT_ID)
	districtResponse := controller.DistrictService.FindById(httpRequest.Context(), districtId)
	helper.WriteSuccessResponse(writer, districtResponse)
}

func (controller *DistrictControllerImpl) FindAllWithPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := make(map[string]any)
	searchBy["name"] = helper.GetQueryParam(httpRequest, DISTRICT_NAME)
	searchBy["cityId"] = helper.StringToInt64(helper.GetQueryParam(httpRequest, CITY_ID))
	searchBy["page"] = helper.GetPageOrSize(httpRequest, constant.PAGE)
	searchBy["size"] = helper.GetPageOrSize(httpRequest, constant.SIZE)

	districtResponses := controller.DistrictService.FindAllWithPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, districtResponses)
}

func (controller *DistrictControllerImpl) FindAllWithoutPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := make(map[string]any)
	searchBy["name"] = helper.GetQueryParam(httpRequest, DISTRICT_NAME)
	searchBy["cityId"] = getCityId(helper.GetQueryParam(httpRequest, CITY_ID))

	districtResponses := controller.DistrictService.FindAllWithoutPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, districtResponses)
}

func getCityId(cityId string) int64 {
	if len(cityId) == 0 {
		exception.PanicErrorBadRequest(errors.New(constant.CITY_ID_REQUIRED))
	}
	return helper.StringToInt64(cityId)
}
