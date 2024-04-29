package helper

import (
	"minang-kos-service/constant"
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

func GetPageOrSize(httpRequest *http.Request, name string) int {
	value := GetQueryParam(httpRequest, name)
	if len(value) == 0 {
		switch name {
		case constant.PAGE:
			return 1
		case constant.SIZE:
			return 10
		default:
			return 1
		}
	}
	return StringToInt(value)
}
