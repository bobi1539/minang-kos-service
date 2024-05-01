package endpoint

import (
	"minang-kos-service/controller"

	"github.com/julienschmidt/httprouter"
)

const VILLAGES = "/api/villages"
const VILLAGES_ALL = VILLAGES + "/all"
const VILLAGES_VILLAGE = VILLAGES + "/village/:villageId"

func SetVillageEndpoint(router *httprouter.Router, villageController controller.VillageController) {
	router.POST(VILLAGES, villageController.Create)
	router.PUT(VILLAGES_VILLAGE, villageController.Update)
	router.GET(VILLAGES_ALL, villageController.FindAllWithoutPagination)
	router.GET(VILLAGES_VILLAGE, villageController.FindById)
	router.GET(VILLAGES, villageController.FindAllWithPagination)
	router.DELETE(VILLAGES_VILLAGE, villageController.Delete)
}
