package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CrudController interface {
	Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
}
