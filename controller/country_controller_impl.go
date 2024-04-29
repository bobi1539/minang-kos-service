package controller

import (
	"minang-kos-service/constant"
	"minang-kos-service/helper"
	"minang-kos-service/model/web/request"
	"minang-kos-service/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const COUNTRY_ID = "countryId"

type CountryControllerImpl struct {
	CountryService service.CountryService
}

func NewCountryController(countryService service.CountryService) CountryController {
	return &CountryControllerImpl{
		CountryService: countryService,
	}
}

func (controller *CountryControllerImpl) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	countryRequest := request.CountryCreateRequest{}
	helper.ReadFromRequestBody(httpRequest, &countryRequest)

	countryResponse := controller.CountryService.Create(httpRequest.Context(), countryRequest)
	helper.WriteSuccessResponse(writer, countryResponse)
}

func (controller *CountryControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	countryRequest := request.CountryUpdateRequest{}
	helper.ReadFromRequestBody(httpRequest, &countryRequest)

	countryRequest.Id = helper.GetIdFromPath(params, COUNTRY_ID)
	countryResponse := controller.CountryService.Update(httpRequest.Context(), countryRequest)
	helper.WriteSuccessResponse(writer, countryResponse)
}

func (controller *CountryControllerImpl) Delete(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	countryId := helper.GetIdFromPath(params, COUNTRY_ID)
	controller.CountryService.Delete(httpRequest.Context(), countryId)
	helper.WriteSuccessResponse(writer, nil)
}

func (controller *CountryControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	countryId := helper.GetIdFromPath(params, COUNTRY_ID)
	countryResponse := controller.CountryService.FindById(httpRequest.Context(), countryId)
	helper.WriteSuccessResponse(writer, countryResponse)
}

func (controller *CountryControllerImpl) FindAllWithPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := make(map[string]any)
	searchBy["name"] = helper.GetQueryParam(httpRequest, "name")
	searchBy["page"] = helper.GetPageOrSize(httpRequest, constant.PAGE)
	searchBy["size"] = helper.GetPageOrSize(httpRequest, constant.SIZE)

	countryResponses := controller.CountryService.FindAllWithPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, countryResponses)
}

func (controller *CountryControllerImpl) FindAllWithoutPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := make(map[string]any)
	searchBy["name"] = helper.GetQueryParam(httpRequest, "name")

	countryResponses := controller.CountryService.FindAllWithoutPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, countryResponses)
}
