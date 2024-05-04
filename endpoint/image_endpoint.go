package endpoint

import (
	"minang-kos-service/controller"

	"github.com/julienschmidt/httprouter"
)

const IMAGES = "/api/static/images"
const IMAGES_KOS_BEDROOM = IMAGES + "/kos-bedroom"

func SetImageEndpoint(router *httprouter.Router) {
	imageController := getImageController()

	router.GET(IMAGES_KOS_BEDROOM, imageController.FindImageKosBedroom)
}

func getImageController() controller.ImageController {
	return controller.NewImageController()
}
