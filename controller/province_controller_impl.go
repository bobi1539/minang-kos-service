package controller

import (
	"minang-kos-service/helper"
	"minang-kos-service/model/web/request"
	"minang-kos-service/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const PROVINCE_ID = "provinceId"

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
	panic("imp")
}

func (controller *ProvinceControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	provinceId := helper.GetIdFromPath(params, PROVINCE_ID)
	provinceResponse := controller.ProvinceService.FindById(httpRequest.Context(), provinceId)
	helper.WriteSuccessResponse(writer, provinceResponse)
}

func (controller *ProvinceControllerImpl) FindAllWithPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	panic("imp")
}

func (controller *ProvinceControllerImpl) FindAllWithoutPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	panic("imp")
}
