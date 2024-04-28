package controller

import (
	"minang-kos-service/helper"
	"minang-kos-service/model/web/request"
	"minang-kos-service/model/web/response"
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

	countryResponse := controller.CountryService.Create(httpRequest.Context(), countryCreateRequest).(response.CountryResponse)
	webResponse := helper.BuildSuccessResponse(countryResponse)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CountryControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	countryUpdateRequest := request.CountryUpdateRequest{}
	helper.ReadFromRequestBody(httpRequest, &countryUpdateRequest)

	countryId := getCountryIdFromPath(params)
	countryUpdateRequest.Id = countryId

	countryResponse := controller.CountryService.Update(httpRequest.Context(), countryUpdateRequest).(response.CountryResponse)
	webResponse := helper.BuildSuccessResponse(countryResponse)
	helper.WriteToResponseBody(writer, webResponse)
}

func getCountryIdFromPath(params httprouter.Params) int64 {
	countryId := params.ByName("countryId")
	id, err := strconv.Atoi(countryId)
	helper.PanicIfError(err)
	return int64(id)
}
