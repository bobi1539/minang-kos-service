package endpoint

import (
	"database/sql"
	"minang-kos-service/controller"
	"minang-kos-service/repository"
	"minang-kos-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

const ROLES = "/api/roles"
const ROLES_ALL = ROLES + "/all"
const ROLES_ROLE = ROLES + "/role/:roleId"

func SetRoleEndpoint(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	roleController := getRoleController(db, validate)

	router.POST(ROLES, roleController.Create)
	router.PUT(ROLES_ROLE, roleController.Update)
	router.GET(ROLES_ALL, roleController.FindAllWithoutPagination)
	router.GET(ROLES_ROLE, roleController.FindById)
	router.GET(ROLES, roleController.FindAllWithPagination)
	router.DELETE(ROLES_ROLE, roleController.Delete)
}

func getRoleController(db *sql.DB, validate *validator.Validate) controller.CountryController {
	roleService := service.NewRoleService(getRoleRepository(), db, validate)
	return controller.NewRoleController(roleService)
}

func getRoleRepository() repository.RoleRepository {
	return repository.NewRoleRepository()
}
