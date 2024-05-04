package controller

import (
	"minang-kos-service/helper"
	"minang-kos-service/model/web/request"
	"minang-kos-service/model/web/search"
	"minang-kos-service/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const KOS_TYPE_ID = "kosTypeId"
const KOS_TYPE_NAME = "name"

type KosTypeControllerImpl struct {
	KosTypeService service.KosTypeService
}

func NewKosTypeController(kosTypeService service.KosTypeService) KosTypeController {
	return &KosTypeControllerImpl{
		KosTypeService: kosTypeService,
	}
}

func (controller *KosTypeControllerImpl) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	kosTypeRequest := request.KosTypeCreateRequest{}
	helper.ReadFromRequestBody(httpRequest, &kosTypeRequest)

	kosTypeResponse := controller.KosTypeService.Create(httpRequest.Context(), kosTypeRequest)
	helper.WriteSuccessResponse(writer, kosTypeResponse)
}

func (controller *KosTypeControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	kosTypeRequest := request.KosTypeUpdateRequest{}
	helper.ReadFromRequestBody(httpRequest, &kosTypeRequest)

	kosTypeRequest.Id = helper.GetIdFromPath(params, KOS_TYPE_ID)
	kosTypeResponse := controller.KosTypeService.Update(httpRequest.Context(), kosTypeRequest)
	helper.WriteSuccessResponse(writer, kosTypeResponse)
}

func (controller *KosTypeControllerImpl) Delete(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	kosTypeId := helper.GetIdFromPath(params, KOS_TYPE_ID)
	controller.KosTypeService.Delete(httpRequest.Context(), kosTypeId)
	helper.WriteSuccessResponse(writer, nil)
}

func (controller *KosTypeControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	kosTypeId := helper.GetIdFromPath(params, KOS_TYPE_ID)
	kosTypeResponse := controller.KosTypeService.FindById(httpRequest.Context(), kosTypeId)
	helper.WriteSuccessResponse(writer, kosTypeResponse)
}

func (controller *KosTypeControllerImpl) FindAllWithPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := search.KosTypeSearch{
		Name:     helper.GetQueryParam(httpRequest, KOS_TYPE_NAME),
		PageSize: search.BuildPageSizeFromRequest(httpRequest),
	}

	kosTypeResponses := controller.KosTypeService.FindAllWithPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, kosTypeResponses)
}

func (controller *KosTypeControllerImpl) FindAllWithoutPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := search.KosTypeSearch{
		Name:     helper.GetQueryParam(httpRequest, KOS_TYPE_NAME),
		PageSize: search.BuildPageSize(0, 0),
	}

	kosTypeResponses := controller.KosTypeService.FindAllWithoutPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, kosTypeResponses)
}
