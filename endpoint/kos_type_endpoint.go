package endpoint

import (
	"minang-kos-service/controller"

	"github.com/julienschmidt/httprouter"
)

const KOS_TYPES = "/api/kos-types"
const KOS_TYPES_ALL = KOS_TYPES + "/all"
const KOS_TYPES_KOS_TYPE = KOS_TYPES + "/kos-type/:kosTypeId"

func SetKosTypeEndpoint(router *httprouter.Router, kosTypeController controller.KosTypeController) {
	router.POST(KOS_TYPES, kosTypeController.Create)
	router.PUT(KOS_TYPES_KOS_TYPE, kosTypeController.Update)
	router.GET(KOS_TYPES_ALL, kosTypeController.FindAllWithoutPagination)
	router.GET(KOS_TYPES_KOS_TYPE, kosTypeController.FindById)
	router.GET(KOS_TYPES, kosTypeController.FindAllWithPagination)
	router.DELETE(KOS_TYPES_KOS_TYPE, kosTypeController.Delete)
}
