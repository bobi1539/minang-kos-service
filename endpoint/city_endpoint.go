package endpoint

import (
	"database/sql"
	"minang-kos-service/controller"
	"minang-kos-service/repository"
	"minang-kos-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

const CITIES = "/api/cities"
const CITIES_ALL = CITIES + "/all"
const CITIES_CITY = CITIES + "/city/:cityId"

func SetCityEndpoint(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	cityController := getCityController(db, validate)

	router.POST(CITIES, cityController.Create)
	router.PUT(CITIES_CITY, cityController.Update)
	router.GET(CITIES_ALL, cityController.FindAllWithoutPagination)
	router.GET(CITIES_CITY, cityController.FindById)
	router.GET(CITIES, cityController.FindAllWithPagination)
	router.DELETE(CITIES_CITY, cityController.Delete)
}

func getCityController(db *sql.DB, validate *validator.Validate) controller.CityController {
	cityService := service.NewCityService(getCityRepository(), getProvinceRepository(), db, validate)
	return controller.NewCityController(cityService)
}

func getCityRepository() repository.CityRepository {
	return repository.NewCityRepository()
}
