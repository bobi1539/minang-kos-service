package endpoint

import (
	"minang-kos-service/controller"

	"github.com/julienschmidt/httprouter"
)

const DISTRICTS = "/api/districts"
const DISTRICTS_ALL = DISTRICTS + "/all"
const DISTRICTS_DISTRICT = DISTRICTS + "/district/:districtId"

func SetDistrictEndpoint(router *httprouter.Router, districtController controller.DistrictController) {
	router.POST(DISTRICTS, districtController.Create)
	router.PUT(DISTRICTS_DISTRICT, districtController.Update)
	router.GET(DISTRICTS_ALL, districtController.FindAllWithoutPagination)
	router.GET(DISTRICTS_DISTRICT, districtController.FindById)
	router.GET(DISTRICTS, districtController.FindAllWithPagination)
	router.DELETE(DISTRICTS_DISTRICT, districtController.Delete)
}
