package controller

import (
	"minang-kos-service/helper"
	"minang-kos-service/model/web/request"
	"minang-kos-service/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	loginRequest := request.AuthLoginRequest{}
	helper.ReadFromRequestBody(httpRequest, &loginRequest)

	loginResponse := controller.AuthService.Login(httpRequest.Context(), loginRequest)
	helper.WriteSuccessResponse(writer, loginResponse)
}
