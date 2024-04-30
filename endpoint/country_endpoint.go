package endpoint

import (
	"minang-kos-service/controller"

	"github.com/julienschmidt/httprouter"
)

const COUNTRIES = "/api/countries"
const COUNTRIES_ALL = COUNTRIES + "/all"
const COUNTRIES_COUNTRY = COUNTRIES + "/country/:countryId"

func SetCountryEndpoint(router *httprouter.Router, countryController controller.CountryController) {
	router.POST(COUNTRIES, countryController.Create)
	router.PUT(COUNTRIES_COUNTRY, countryController.Update)
	router.GET(COUNTRIES_ALL, countryController.FindAllWithoutPagination)
	router.GET(COUNTRIES_COUNTRY, countryController.FindById)
	router.GET(COUNTRIES, countryController.FindAllWithPagination)
	router.DELETE(COUNTRIES_COUNTRY, countryController.Delete)
}
