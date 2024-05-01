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

const VILLAGE_ID = "villageId"
const VILLAGE_NAME = "name"

type VillageControllerImpl struct {
	VillageService service.VillageService
}

func NewVillageController(villageService service.VillageService) VillageController {
	return &VillageControllerImpl{
		VillageService: villageService,
	}
}

func (controller *VillageControllerImpl) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	villageRequest := request.VillageCreateRequest{}
	helper.ReadFromRequestBody(httpRequest, &villageRequest)

	villageResponse := controller.VillageService.Create(httpRequest.Context(), villageRequest)
	helper.WriteSuccessResponse(writer, villageResponse)
}
func (controller *VillageControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	villageRequest := request.VillageUpdateRequest{}
	helper.ReadFromRequestBody(httpRequest, &villageRequest)

	villageRequest.Id = helper.GetIdFromPath(params, VILLAGE_ID)
	villageResponse := controller.VillageService.Update(httpRequest.Context(), villageRequest)
	helper.WriteSuccessResponse(writer, villageResponse)
}
func (controller *VillageControllerImpl) Delete(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	villageId := helper.GetIdFromPath(params, VILLAGE_ID)
	controller.VillageService.Delete(httpRequest.Context(), villageId)
	helper.WriteSuccessResponse(writer, nil)
}
func (controller *VillageControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	villageId := helper.GetIdFromPath(params, VILLAGE_ID)
	villageResponse := controller.VillageService.FindById(httpRequest.Context(), villageId)
	helper.WriteSuccessResponse(writer, villageResponse)
}
func (controller *VillageControllerImpl) FindAllWithPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := make(map[string]any)
	searchBy["name"] = helper.GetQueryParam(httpRequest, VILLAGE_NAME)
	searchBy["districtId"] = helper.StringToInt64(helper.GetQueryParam(httpRequest, DISTRICT_ID))
	searchBy["page"] = helper.GetPageOrSize(httpRequest, constant.PAGE)
	searchBy["size"] = helper.GetPageOrSize(httpRequest, constant.SIZE)

	villageResponses := controller.VillageService.FindAllWithPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, villageResponses)
}
func (controller *VillageControllerImpl) FindAllWithoutPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := make(map[string]any)
	searchBy["name"] = helper.GetQueryParam(httpRequest, VILLAGE_NAME)
	searchBy["districtId"] = getDistrictId(helper.GetQueryParam(httpRequest, DISTRICT_ID))

	villageResponses := controller.VillageService.FindAllWithoutPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, villageResponses)
}

func getDistrictId(districtId string) int64 {
	if len(districtId) == 0 {
		exception.PanicErrorBadRequest(errors.New(constant.DISTRICT_ID_REQUIRED))
	}
	return helper.StringToInt64(districtId)
}
