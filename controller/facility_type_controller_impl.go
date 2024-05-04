package controller

import (
	"minang-kos-service/helper"
	"minang-kos-service/model/web/request"
	"minang-kos-service/model/web/search"
	"minang-kos-service/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const FACILITY_TYPE_ID = "facilityTypeId"
const FACILITY_TYPE_NAME = "name"

type FacilityTypeControllerImpl struct {
	FacilitTypeService service.FacilityTypeService
}

func NewFacilityTypeController(facilityTypeService service.FacilityTypeService) FacilityTypeController {
	return &FacilityTypeControllerImpl{
		FacilitTypeService: facilityTypeService,
	}
}

func (controller *FacilityTypeControllerImpl) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	facilityTypeRequest := request.FacilityTypeCreateRequest{}
	helper.ReadFromRequestBody(httpRequest, &facilityTypeRequest)

	facilityTypeResponse := controller.FacilitTypeService.Create(httpRequest.Context(), facilityTypeRequest)
	helper.WriteSuccessResponse(writer, facilityTypeResponse)
}

func (controller *FacilityTypeControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	facilityTypeRequest := request.FacilityTypeUpdateRequest{}
	helper.ReadFromRequestBody(httpRequest, &facilityTypeRequest)

	facilityTypeRequest.Id = helper.GetIdFromPath(params, FACILITY_TYPE_ID)
	facilityTypeResponse := controller.FacilitTypeService.Update(httpRequest.Context(), facilityTypeRequest)
	helper.WriteSuccessResponse(writer, facilityTypeResponse)
}

func (controller *FacilityTypeControllerImpl) Delete(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	facilityTypeId := helper.GetIdFromPath(params, FACILITY_TYPE_ID)
	controller.FacilitTypeService.Delete(httpRequest.Context(), facilityTypeId)
	helper.WriteSuccessResponse(writer, nil)
}

func (controller *FacilityTypeControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	facilityTypeId := helper.GetIdFromPath(params, FACILITY_TYPE_ID)
	facilityTypeResponse := controller.FacilitTypeService.FindById(httpRequest.Context(), facilityTypeId)
	helper.WriteSuccessResponse(writer, facilityTypeResponse)
}

func (controller *FacilityTypeControllerImpl) FindAllWithPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := search.FacilityTypeSearch{
		Name:     helper.GetQueryParam(httpRequest, FACILITY_TYPE_NAME),
		PageSize: search.BuildPageSizeFromRequest(httpRequest),
	}

	facilityTypeResponses := controller.FacilitTypeService.FindAllWithPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, facilityTypeResponses)
}

func (controller *FacilityTypeControllerImpl) FindAllWithoutPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := search.FacilityTypeSearch{
		Name:     helper.GetQueryParam(httpRequest, FACILITY_TYPE_NAME),
		PageSize: search.BuildPageSize(0, 0),
	}

	facilityTypeResponses := controller.FacilitTypeService.FindAllWithoutPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, facilityTypeResponses)
}
