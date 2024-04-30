package endpoint

import (
	"minang-kos-service/controller"

	"github.com/julienschmidt/httprouter"
)

const AUTH = "/api/auth"
const AUTH_LOGIN = AUTH + "/login"

func SetAuthEndpoint(router *httprouter.Router, authController controller.AuthController) {
	router.POST(AUTH_LOGIN, authController.Login)
}
