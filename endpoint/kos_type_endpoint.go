package endpoint

import (
	"database/sql"
	"minang-kos-service/controller"
	"minang-kos-service/repository"
	"minang-kos-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

const KOS_TYPES = "/api/kos-types"
const KOS_TYPES_ALL = KOS_TYPES + "/all"
const KOS_TYPES_KOS_TYPE = KOS_TYPES + "/kos-type/:kosTypeId"

func SetKosTypeEndpoint(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	kosTypeController := getKosTypeController(db, validate)

	router.POST(KOS_TYPES, kosTypeController.Create)
	router.PUT(KOS_TYPES_KOS_TYPE, kosTypeController.Update)
	router.GET(KOS_TYPES_ALL, kosTypeController.FindAllWithoutPagination)
	router.GET(KOS_TYPES_KOS_TYPE, kosTypeController.FindById)
	router.GET(KOS_TYPES, kosTypeController.FindAllWithPagination)
	router.DELETE(KOS_TYPES_KOS_TYPE, kosTypeController.Delete)
}

func getKosTypeController(db *sql.DB, validate *validator.Validate) controller.KosTypeController {
	kosTypeService := service.NewKosTypeService(getKosTypeRepository(), db, validate)
	return controller.NewKosTypeController(kosTypeService)
}

func getKosTypeRepository() repository.KosTypeRepository {
	return repository.NewKosTypeRepository()
}
