package exception

import (
	"fmt"
	"minang-kos-service/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {
	fmt.Println(err)
	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request)
}

func internalServerError(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := helper.BuildInternalServerErrorResponse()
	helper.WriteToResponseBody(writer, webResponse)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err any) bool {
	_, ok := err.(ErrorNotFound)
	if ok {
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := helper.BuildNotFoundErrorResponse()
		helper.WriteToResponseBody(writer, webResponse)
		return true
	}
	return false
}

func validationError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := helper.BuildValidationErrorResponse(exception.Error())
		helper.WriteToResponseBody(writer, webResponse)
		return true
	}
	return false
}

func PanicErrorNotFound(err error) {
	if err != nil {
		panic(NewErrorNotFound(err.Error()))
	}
}
