package endpoint

import (
	"database/sql"
	"minang-kos-service/controller"
	"minang-kos-service/repository"
	"minang-kos-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

const DISTRICTS = "/api/districts"
const DISTRICTS_ALL = DISTRICTS + "/all"
const DISTRICTS_DISTRICT = DISTRICTS + "/district/:districtId"

func SetDistrictEndpoint(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	districtController := getDistrictController(db, validate)

	router.POST(DISTRICTS, districtController.Create)
	router.PUT(DISTRICTS_DISTRICT, districtController.Update)
	router.GET(DISTRICTS_ALL, districtController.FindAllWithoutPagination)
	router.GET(DISTRICTS_DISTRICT, districtController.FindById)
	router.GET(DISTRICTS, districtController.FindAllWithPagination)
	router.DELETE(DISTRICTS_DISTRICT, districtController.Delete)
}

func getDistrictController(db *sql.DB, validate *validator.Validate) controller.DistrictController {
	districtService := service.NewDistrictService(getDistrictRepository(), getCityRepository(), db, validate)
	return controller.NewDistrictController(districtService)
}

func getDistrictRepository() repository.DistrictRepository {
	return repository.NewDistrictRepository()
}
