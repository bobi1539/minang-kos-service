package controller

import (
	"minang-kos-service/helper"
	"minang-kos-service/model/web/request"
	"minang-kos-service/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CountryControllerImpl struct {
	CountryService service.CountryService
}

func NewCountryController(countryService service.CountryService) CountryController {
	return &CountryControllerImpl{
		CountryService: countryService,
	}
}

func (controller *CountryControllerImpl) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	countryCreateRequest := request.CountryCreateRequest{}
	helper.ReadFromRequestBody(httpRequest, &countryCreateRequest)

	countryResponse := controller.CountryService.Create(httpRequest.Context(), countryCreateRequest)
	writeCountryResponseToResponseBody(writer, countryResponse)
}

func (controller *CountryControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	countryUpdateRequest := request.CountryUpdateRequest{}
	helper.ReadFromRequestBody(httpRequest, &countryUpdateRequest)

	countryId := getCountryIdFromPath(params)
	countryUpdateRequest.Id = countryId

	countryResponse := controller.CountryService.Update(httpRequest.Context(), countryUpdateRequest)
	writeCountryResponseToResponseBody(writer, countryResponse)
}

func (controller *CountryControllerImpl) Delete(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {

}

func (controller *CountryControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	countryId := getCountryIdFromPath(params)
	countryResponse := controller.CountryService.FindById(httpRequest.Context(), countryId)
	writeCountryResponseToResponseBody(writer, countryResponse)
}

func getCountryIdFromPath(params httprouter.Params) int64 {
	countryId := params.ByName("countryId")
	id, err := strconv.Atoi(countryId)
	helper.PanicIfError(err)
	return int64(id)
}

func writeCountryResponseToResponseBody(writer http.ResponseWriter, countryResponse any) {
	webResponse := helper.BuildSuccessResponse(countryResponse)
	helper.WriteToResponseBody(writer, webResponse)
}
