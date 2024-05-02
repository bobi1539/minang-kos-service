package endpoint

import (
	"minang-kos-service/controller"

	"github.com/julienschmidt/httprouter"
)

const FACILITY_TYPES = "/api/facility-types"
const FACILITY_TYPES_ALL = FACILITY_TYPES + "/all"
const FACILITY_TYPES_FACILITY_TYPE = FACILITY_TYPES + "/facility-type/:facilityTypeId"

func SetFacilityTypeEndpoint(router *httprouter.Router, facilityTypeController controller.FacilityTypeController) {
	router.POST(FACILITY_TYPES, facilityTypeController.Create)
	router.PUT(FACILITY_TYPES_FACILITY_TYPE, facilityTypeController.Update)
	router.GET(FACILITY_TYPES_ALL, facilityTypeController.FindAllWithoutPagination)
	router.GET(FACILITY_TYPES_FACILITY_TYPE, facilityTypeController.FindById)
	router.GET(FACILITY_TYPES, facilityTypeController.FindAllWithPagination)
	router.DELETE(FACILITY_TYPES_FACILITY_TYPE, facilityTypeController.Delete)
}
