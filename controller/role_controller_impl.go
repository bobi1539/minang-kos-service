package controller

import (
	"minang-kos-service/constant"
	"minang-kos-service/helper"
	"minang-kos-service/model/web/request"
	"minang-kos-service/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const ROLE_ID = "roleId"

type RoleControllerImpl struct {
	RoleService service.RoleService
}

func NewRoleController(roleService service.RoleService) RoleController {
	return &RoleControllerImpl{
		RoleService: roleService,
	}
}

func (controller *RoleControllerImpl) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	roleRequest := request.RoleCreateRequest{}
	helper.ReadFromRequestBody(httpRequest, &roleRequest)

	roleResponse := controller.RoleService.Create(httpRequest.Context(), roleRequest)
	helper.WriteSuccessResponse(writer, roleResponse)
}

func (controller *RoleControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	roleRequest := request.RoleUpdateRequest{}
	helper.ReadFromRequestBody(httpRequest, &roleRequest)

	roleRequest.Id = helper.GetIdFromPath(params, ROLE_ID)
	roleResponse := controller.RoleService.Update(httpRequest.Context(), roleRequest)
	helper.WriteSuccessResponse(writer, roleResponse)
}

func (controller *RoleControllerImpl) Delete(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	roleId := helper.GetIdFromPath(params, ROLE_ID)
	controller.RoleService.Delete(httpRequest.Context(), roleId)
	helper.WriteSuccessResponse(writer, nil)
}

func (controller *RoleControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	roleId := helper.GetIdFromPath(params, ROLE_ID)
	roleResponse := controller.RoleService.FindById(httpRequest.Context(), roleId)
	helper.WriteSuccessResponse(writer, roleResponse)
}

func (controller *RoleControllerImpl) FindAllWithPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := make(map[string]any)
	searchBy["name"] = helper.GetQueryParam(httpRequest, "name")
	searchBy["page"] = helper.GetPageOrSize(httpRequest, constant.PAGE)
	searchBy["size"] = helper.GetPageOrSize(httpRequest, constant.SIZE)

	roleResponses := controller.RoleService.FindAllWithPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, roleResponses)
}

func (controller *RoleControllerImpl) FindAllWithoutPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	searchBy := make(map[string]any)
	searchBy["name"] = helper.GetQueryParam(httpRequest, "name")

	roleResponses := controller.RoleService.FindAllWithoutPagination(httpRequest.Context(), searchBy)
	helper.WriteSuccessResponse(writer, roleResponses)
}
