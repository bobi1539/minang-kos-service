package controller

import (
	"minang-kos-service/helper"
	"minang-kos-service/model/web/request"
	"minang-kos-service/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const KOS_BEDROOM_ID = "kosBedroomId"

type KosBedroomControllerImpl struct {
	KosBedroomService service.KosBedroomService
}

func NewKosBedroomController(kosBedroomService service.KosBedroomService) KosBedroomController {
	return &KosBedroomControllerImpl{
		KosBedroomService: kosBedroomService,
	}
}

func (controller *KosBedroomControllerImpl) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	kosBedroomRequest := request.KosBedroomCreateRequest{}
	helper.ReadFromRequestBody(httpRequest, &kosBedroomRequest)

	kosBedroomResponse := controller.KosBedroomService.Create(httpRequest.Context(), kosBedroomRequest)
	helper.WriteSuccessResponse(writer, kosBedroomResponse)
}

func (controller *KosBedroomControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	panic("imp")
}

func (controller *KosBedroomControllerImpl) Delete(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	panic("imp")
}

func (controller *KosBedroomControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	kosBedroomId := helper.GetIdFromPath(params, KOS_BEDROOM_ID)
	kosBedroomResponse := controller.KosBedroomService.FindById(httpRequest.Context(), kosBedroomId)
	helper.WriteSuccessResponse(writer, kosBedroomResponse)
}

func (controller *KosBedroomControllerImpl) FindAllWithPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	panic("imp")
}

func (controller *KosBedroomControllerImpl) FindAllWithoutPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	panic("imp")
}
