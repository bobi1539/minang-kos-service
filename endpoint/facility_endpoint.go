package endpoint

import (
	"database/sql"
	"minang-kos-service/controller"
	"minang-kos-service/repository"
	"minang-kos-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

const FACILITIES = "/api/facilities"
const FACILITIES_ALL = FACILITIES + "/all"
const FACILITIES_FACILITY = FACILITIES + "/facility/:facilityId"

func SetFacilityEndpoint(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	facilityController := getFacilityController(db, validate)

	router.POST(FACILITIES, facilityController.Create)
	router.PUT(FACILITIES_FACILITY, facilityController.Update)
	router.GET(FACILITIES_ALL, facilityController.FindAllWithoutPagination)
	router.GET(FACILITIES_FACILITY, facilityController.FindById)
	router.GET(FACILITIES, facilityController.FindAllWithPagination)
	router.DELETE(FACILITIES_FACILITY, facilityController.Delete)
}

func getFacilityController(db *sql.DB, validate *validator.Validate) controller.FacilityController {
	facilityService := service.NewFacilityService(getFacilityRepository(), getFacilityTypeRepository(), db, validate)
	return controller.NewFacilityController(facilityService)
}

func getFacilityRepository() repository.FacilityRepository {
	return repository.NewFacilityRepository()
}
