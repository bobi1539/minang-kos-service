package controller

import (
	"minang-kos-service/constant"
	"minang-kos-service/helper"
	"minang-kos-service/model/web/request"
	"minang-kos-service/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const FACILITY_ID = "facilityId"
const FACILITY_NAME = "name"

type FacilityControllerImpl struct {
	FacilityService service.FacilityService
}

func NewFacilityController(facilityService service.FacilityService) FacilityController {
	return &FacilityControllerImpl{
		FacilityService: facilityService,
	}
}

func (controller *FacilityControllerImpl) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	facilityRequest := request.FacilityCreateRequest{}
	helper.ReadFromRequestBody(httpRequest, &facilityRequest)

	facilityResponse := controller.FacilityService.Create(httpRequest.Context(), facilityRequest)
	helper.WriteSuccessResponse(writer, facilityResponse)
}

func (controller *FacilityControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	facilityRequest := request.FacilityUpdateRequest{}
	helper.ReadFromRequestBody(httpRequest, &facilityRequest)

	facilityRequest.Id = helper.GetIdFromPath(params, FACILITY_ID)
	facilityResponse := controller.FacilityService.Update(httpRequest.Context(), facilityRequest)
	helper.WriteSuccessResponse(writer, facilityResponse)
}

func (controller *FacilityControllerImpl) Delete(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	facilityId := helper.GetIdFromPath(params, FACILITY_ID)
	controller.FacilityService.Delete(httpRequest.Context(), facilityId)
	helper.WriteSuccessResponse(writer, nil)
}

func (controller *FacilityControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	facilityId := helper.GetIdFromPath(params, FACILITY_ID)
	facilityResponse := controller.FacilityService.FindById(httpRequest.Context(), facilityId)
	helper.WriteSuccessResponse(writer, facilityResponse)
}

func (controller *FacilityControllerImpl) FindAllWithPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := make(map[string]any)
	searchBy["name"] = helper.GetQueryParam(httpRequest, FACILITY_NAME)
	searchBy["facilityTypeId"] = helper.StringToInt64(helper.GetQueryParam(httpRequest, FACILITY_TYPE_ID))
	searchBy["page"] = helper.GetPageOrSize(httpRequest, constant.PAGE)
	searchBy["size"] = helper.GetPageOrSize(httpRequest, constant.SIZE)

	facilityResponses := controller.FacilityService.FindAllWithPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, facilityResponses)
}

func (controller *FacilityControllerImpl) FindAllWithoutPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := make(map[string]any)
	searchBy["name"] = helper.GetQueryParam(httpRequest, FACILITY_NAME)
	searchBy["facilityTypeId"] = helper.StringToInt64(helper.GetQueryParam(httpRequest, FACILITY_TYPE_ID))

	facilityResponses := controller.FacilityService.FindAllWithoutPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, facilityResponses)
}
