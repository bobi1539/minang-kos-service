package endpoint

import (
	"minang-kos-service/controller"

	"github.com/julienschmidt/httprouter"
)

const FACILITIES = "/api/facilities"
const FACILITIES_ALL = FACILITIES + "/all"
const FACILITIES_FACILITY = FACILITIES + "/facility/:facilityId"

func SetFacilityEndpoint(router *httprouter.Router, facilityController controller.FacilityController) {
	router.POST(FACILITIES, facilityController.Create)
	router.PUT(FACILITIES_FACILITY, facilityController.Update)
	router.GET(FACILITIES_ALL, facilityController.FindAllWithoutPagination)
	router.GET(FACILITIES_FACILITY, facilityController.FindById)
	router.GET(FACILITIES, facilityController.FindAllWithPagination)
	router.DELETE(FACILITIES_FACILITY, facilityController.Delete)
}
