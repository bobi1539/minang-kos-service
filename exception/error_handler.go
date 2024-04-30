package exception

import (
	"minang-kos-service/helper"
	"net/http"
	"runtime/debug"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger = helper.GetLogger()

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {
	log.Error(err, string(debug.Stack()))

	if badRequestError(writer, request, err) {
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

func badRequestError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(ErrorBadRequest)
	if ok {
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := helper.BuildBadRequestErrorResponse(exception.Error)
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

func PanicErrorBadRequest(err error) {
	if err != nil {
		panic(NewErrorBadRequest(err.Error()))
	}
}
