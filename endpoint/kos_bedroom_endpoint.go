package endpoint

import (
	"minang-kos-service/controller"

	"github.com/julienschmidt/httprouter"
)

const KOS_BEDROOMS = "/api/kos-bedrooms"
const KOS_BEDROOMS_ALL = KOS_BEDROOMS + "/all"
const KOS_BEDROOMS_KOS_BEDROOM = KOS_BEDROOMS + "/kos-bedroom/:kodBedroomId"

func SetKosBedroomEndpoint(router *httprouter.Router, kosBedroomController controller.KosBedroomController) {
	router.POST(KOS_BEDROOMS, kosBedroomController.Create)
	router.PUT(KOS_BEDROOMS_KOS_BEDROOM, kosBedroomController.Update)
	router.GET(KOS_BEDROOMS_ALL, kosBedroomController.FindAllWithoutPagination)
	router.GET(KOS_BEDROOMS_KOS_BEDROOM, kosBedroomController.FindById)
	router.GET(KOS_BEDROOMS, kosBedroomController.FindAllWithPagination)
	router.DELETE(KOS_BEDROOMS_KOS_BEDROOM, kosBedroomController.Delete)
}
