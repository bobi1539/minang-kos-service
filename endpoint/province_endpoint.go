package endpoint

import (
	"minang-kos-service/controller"

	"github.com/julienschmidt/httprouter"
)

const PROVINCES = "/api/provinces"
const PROVINCES_ALL = PROVINCES + "/all"
const PROVINCES_PROVINCE = PROVINCES + "/province/:provinceId"

func SetProvinceEndpoint(router *httprouter.Router, provinceController controller.ProvinceController) {
	router.POST(PROVINCES, provinceController.Create)
	router.PUT(PROVINCES_PROVINCE, provinceController.Update)
	router.GET(PROVINCES_ALL, provinceController.FindAllWithoutPagination)
	router.GET(PROVINCES_PROVINCE, provinceController.FindById)
	router.GET(PROVINCES, provinceController.FindAllWithPagination)
	router.DELETE(PROVINCES_PROVINCE, provinceController.Delete)
}
