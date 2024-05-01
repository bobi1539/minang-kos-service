package endpoint

import (
	"minang-kos-service/controller"

	"github.com/julienschmidt/httprouter"
)

const CITIES = "/api/cities"
const CITIES_ALL = CITIES + "/all"
const CITIES_CITY = CITIES + "/city/:cityId"

func SetCityEndpoint(router *httprouter.Router, cityController controller.CityController) {
	router.POST(CITIES, cityController.Create)
	router.PUT(CITIES_CITY, cityController.Update)
	router.GET(CITIES_ALL, cityController.FindAllWithoutPagination)
	router.GET(CITIES_CITY, cityController.FindById)
	router.GET(CITIES, cityController.FindAllWithPagination)
	router.DELETE(CITIES_CITY, cityController.Delete)
}
