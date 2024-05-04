package endpoint

import (
	"database/sql"
	"minang-kos-service/controller"
	"minang-kos-service/repository"
	"minang-kos-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

const COUNTRIES = "/api/countries"
const COUNTRIES_ALL = COUNTRIES + "/all"
const COUNTRIES_COUNTRY = COUNTRIES + "/country/:countryId"

func SetCountryEndpoint(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	countryController := getCountryController(db, validate)

	router.POST(COUNTRIES, countryController.Create)
	router.PUT(COUNTRIES_COUNTRY, countryController.Update)
	router.GET(COUNTRIES_ALL, countryController.FindAllWithoutPagination)
	router.GET(COUNTRIES_COUNTRY, countryController.FindById)
	router.GET(COUNTRIES, countryController.FindAllWithPagination)
	router.DELETE(COUNTRIES_COUNTRY, countryController.Delete)
}

func getCountryController(db *sql.DB, validate *validator.Validate) controller.CountryController {
	countryService := service.NewCountryService(getCountryRepository(), db, validate)
	return controller.NewCountryController(countryService)
}

func getCountryRepository() repository.CountryRepository {
	return repository.NewCountryRepository()
}
