package endpoint

import (
	"minang-kos-service/controller"

	"github.com/julienschmidt/httprouter"
)

const PROVINCES = "/api/provinces"
const PROVINCES_ALL = PROVINCES + "/all"
const PROVINCES_COUNTRY = PROVINCES + "/province/:provinceId"

func SetProvinceEndpoint(router *httprouter.Router, provinceController controller.ProvinceController) {
	router.POST(PROVINCES, provinceController.Create)
	router.PUT(PROVINCES_COUNTRY, provinceController.Update)
	router.GET(PROVINCES_ALL, provinceController.FindAllWithoutPagination)
	router.GET(PROVINCES_COUNTRY, provinceController.FindById)
	router.GET(PROVINCES, provinceController.FindAllWithPagination)
	router.DELETE(PROVINCES_COUNTRY, provinceController.Delete)
}
