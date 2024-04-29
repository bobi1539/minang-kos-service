package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CrudController interface {
	Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
	FindAllWithPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
	FindAllWithoutPagination(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
}
