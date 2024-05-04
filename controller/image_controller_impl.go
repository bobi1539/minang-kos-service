package controller

import (
	"minang-kos-service/constant"
	"minang-kos-service/helper"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type ImageControllerImpl struct {
}

const FILE_NAME = "filename"

func NewImageController() ImageController {
	return &ImageControllerImpl{}
}

func (controller *ImageControllerImpl) FindImageKosBedroom(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	path := os.Getenv(constant.PATH_FILE_IMAGE_KOS_BEDROOM)
	filename := helper.GetQueryParam(httpRequest, FILE_NAME)

	fullpath := path + filename
	imageFile := helper.ReadFile(fullpath)

	helper.WriteFileResponse(writer, imageFile)
}
