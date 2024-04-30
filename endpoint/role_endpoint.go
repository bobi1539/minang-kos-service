package endpoint

import (
	"minang-kos-service/controller"

	"github.com/julienschmidt/httprouter"
)

const ROLES = "/api/roles"
const ROLES_ALL = ROLES + "/all"
const ROLES_ROLE = ROLES + "/role/:roleId"

func SetRoleEndpoint(router *httprouter.Router, roleController controller.RoleController) {
	router.POST(ROLES, roleController.Create)
	router.PUT(ROLES_ROLE, roleController.Update)
	router.GET(ROLES_ALL, roleController.FindAllWithoutPagination)
	router.GET(ROLES_ROLE, roleController.FindById)
	router.GET(ROLES, roleController.FindAllWithPagination)
	router.DELETE(ROLES_ROLE, roleController.Delete)
}
