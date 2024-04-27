package helper

import "minang-kos-service/model/web/response"

func BuildSuccessResponse(data any) response.WebResponse {
	return response.WebResponse{
		Code:    200,
		Message: "Success | Sukses",
		Data:    data,
	}
}
