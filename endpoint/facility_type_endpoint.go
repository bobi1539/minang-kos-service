package endpoint

import (
	"database/sql"
	"minang-kos-service/controller"
	"minang-kos-service/repository"
	"minang-kos-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

const FACILITY_TYPES = "/api/facility-types"
const FACILITY_TYPES_ALL = FACILITY_TYPES + "/all"
const FACILITY_TYPES_FACILITY_TYPE = FACILITY_TYPES + "/facility-type/:facilityTypeId"

func SetFacilityTypeEndpoint(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	facilityTypeController := getFacilityTypeController(db, validate)

	router.POST(FACILITY_TYPES, facilityTypeController.Create)
	router.PUT(FACILITY_TYPES_FACILITY_TYPE, facilityTypeController.Update)
	router.GET(FACILITY_TYPES_ALL, facilityTypeController.FindAllWithoutPagination)
	router.GET(FACILITY_TYPES_FACILITY_TYPE, facilityTypeController.FindById)
	router.GET(FACILITY_TYPES, facilityTypeController.FindAllWithPagination)
	router.DELETE(FACILITY_TYPES_FACILITY_TYPE, facilityTypeController.Delete)
}

func getFacilityTypeController(db *sql.DB, validate *validator.Validate) controller.FacilityTypeController {
	facilityTypeService := service.NewFacilityTypeService(getFacilityTypeRepository(), db, validate)
	return controller.NewFacilityTypeController(facilityTypeService)
}

func getFacilityTypeRepository() repository.FacilityTypeRepository {
	return repository.NewFacilityTypeRepository()
}
