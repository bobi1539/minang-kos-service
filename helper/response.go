package helper

import (
	"minang-kos-service/constant"
	"minang-kos-service/model/web/response"
	"net/http"
)

func BuildSuccessResponse(data any) response.WebResponse {
	return response.WebResponse{
		Code:    200,
		Message: constant.SUCCESS,
		Data:    data,
	}
}

func BuildInternalServerErrorResponse() response.WebResponse {
	return response.WebResponse{
		Code:    http.StatusInternalServerError,
		Message: constant.INTERNAL_SERVER_ERROR,
	}
}

func BuildNotFoundErrorResponse() response.WebResponse {
	return response.WebResponse{
		Code:    http.StatusBadRequest,
		Message: constant.DATA_NOT_FOUND,
	}
}

func BuildValidationErrorResponse(message string) response.WebResponse {
	return response.WebResponse{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}
