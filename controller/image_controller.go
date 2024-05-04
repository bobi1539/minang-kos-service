package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ImageController interface {
	FindImageKosBedroom(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
}
