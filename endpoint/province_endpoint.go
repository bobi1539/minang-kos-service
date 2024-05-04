package endpoint

import (
	"database/sql"
	"minang-kos-service/controller"
	"minang-kos-service/repository"
	"minang-kos-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

const PROVINCES = "/api/provinces"
const PROVINCES_ALL = PROVINCES + "/all"
const PROVINCES_PROVINCE = PROVINCES + "/province/:provinceId"

func SetProvinceEndpoint(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	provinceController := getProvinceController(db, validate)

	router.POST(PROVINCES, provinceController.Create)
	router.PUT(PROVINCES_PROVINCE, provinceController.Update)
	router.GET(PROVINCES_ALL, provinceController.FindAllWithoutPagination)
	router.GET(PROVINCES_PROVINCE, provinceController.FindById)
	router.GET(PROVINCES, provinceController.FindAllWithPagination)
	router.DELETE(PROVINCES_PROVINCE, provinceController.Delete)
}

func getProvinceController(db *sql.DB, validate *validator.Validate) controller.ProvinceController {
	provinceService := service.NewProvinceService(getProvinceRepository(), getCountryRepository(), db, validate)
	return controller.NewProvinceController(provinceService)
}

func getProvinceRepository() repository.ProvinceRepository {
	return repository.NewProvinceRepository()
}
