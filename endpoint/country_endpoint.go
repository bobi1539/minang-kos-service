package endpoint

import (
	"minang-kos-service/controller"

	"github.com/julienschmidt/httprouter"
)

const COUNTRIES = "/api/countries"
const COUNTRIES_ALL = COUNTRIES + "/all"
const COUNTRIES_COUNTRY = COUNTRIES + "/country/:countryId"

const PROVINCES = "/api/provinces"
const PROVINCES_ALL = PROVINCES + "/all"
const PROVINCES_COUNTRY = PROVINCES + "/province/:provinceId"

func SetCountryEndpoint(router *httprouter.Router, countryController controller.CountryController) {
	router.POST(COUNTRIES, countryController.Create)
	router.PUT(COUNTRIES_COUNTRY, countryController.Update)
	router.GET(COUNTRIES_ALL, countryController.FindAllWithoutPagination)
	router.GET(COUNTRIES_COUNTRY, countryController.FindById)
	router.GET(COUNTRIES, countryController.FindAllWithPagination)
	router.DELETE(COUNTRIES_COUNTRY, countryController.Delete)
}

func SetProvinceEndpoint(router *httprouter.Router, provinceController controller.ProvinceController) {
	router.POST(PROVINCES, provinceController.Create)
	router.PUT(PROVINCES_COUNTRY, provinceController.Update)
	router.GET(PROVINCES_ALL, provinceController.FindAllWithoutPagination)
	router.GET(PROVINCES_COUNTRY, provinceController.FindById)
	router.GET(PROVINCES, provinceController.FindAllWithPagination)
	router.DELETE(PROVINCES_COUNTRY, provinceController.Delete)
}
