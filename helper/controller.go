package helper

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetIdFromPath(params httprouter.Params, name string) int64 {
	countryId := params.ByName(name)
	id, err := strconv.Atoi(countryId)
	PanicIfError(err)
	return int64(id)
}

func GetQueryParam(httpRequest *http.Request, name string) string {
	queryParam := httpRequest.URL.Query()
	return queryParam.Get(name)
}
